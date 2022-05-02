package jwtutils

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/botanist/common-utils/rpc"
	jwtv4 "github.com/cristalhq/jwt/v4"
)

var ErrNoAuthorizationHeader = errors.New("Missing authorization header")
var ErrInvalidAuthorizationHeader = errors.New("The authorization header is invalid")
var ErrSubjectClaimMismatch = errors.New("The registered claim 'sub' doesn't match expected")
var ErrAudienceClaimMismatch = errors.New("The registered claim 'aud' doesn't match expected")

var ErrNotValidAtThisTime = errors.New("The supplied JWT is not valid at this time")

var verifierCache map[string]jwtv4.Verifier
var verifierCacheMutex sync.RWMutex

func init() {
	verifierCache = make(map[string]jwtv4.Verifier)
}

func getVerifier(kid string) (jwtv4.Verifier, error) {
	v, ok := verifierCache[kid]
	if !ok {
		verifierCacheMutex.Lock()
		defer verifierCacheMutex.Unlock()

		v, ok = verifierCache[kid]
		if ok {
			return v, nil
		}

		_, b, err := rpc.Get("auth-service", fmt.Sprintf("/v1/auth/jwt/pubkey/%s", kid), nil, "", nil)
		if err != nil {
			return nil, err
		}

		p, _ := pem.Decode(b)
		if err != nil {
			return nil, err
		}

		pk, err := x509.ParsePKIXPublicKey(p.Bytes)
		if err != nil {
			return nil, err
		}

		v, err = jwtv4.NewVerifierES(jwtv4.ES256, pk.(*ecdsa.PublicKey))
		if err != nil {
			return nil, err
		}

		verifierCache[kid] = v
	}

	return v, nil
}

func VerifyWithClaims(req *http.Request, expectedClaims jwtv4.RegisteredClaims) (*jwtv4.Token, *jwtv4.RegisteredClaims, error) {
	auth := req.Header.Get("Authorization")
	if auth == "" {
		return nil, nil, ErrNoAuthorizationHeader
	}

	if !strings.HasPrefix(auth, "Bearer ") {
		return nil, nil, ErrInvalidAuthorizationHeader
	}

	// Trim Bearer
	auth = auth[len("Bearer "):]

	tok, err := jwtv4.ParseNoVerify([]byte(auth))
	if err != nil {
		return tok, nil, err
	}

	// Verify
	v, err := getVerifier(tok.Header().KeyID)
	if err != nil {
		return nil, nil, err
	}

	err = v.Verify(tok)
	if err != nil {
		return nil, nil, err
	}

	var claims jwtv4.RegisteredClaims
	err = json.Unmarshal(tok.Claims(), &claims)
	if err != nil {
		return tok, &claims, err
	}

	// Check claims
	if !claims.IsValidAt(time.Now()) {
		return tok, &claims, ErrNotValidAtThisTime
	}

	// Check subject
	if expectedClaims.Subject != "" {
		if !claims.IsSubject(expectedClaims.Subject) {
			return tok, &claims, ErrSubjectClaimMismatch
		}
	}

	if expectedClaims.Audience != nil && len(expectedClaims.Audience) > 0 {
		if !claims.IsForAudience(expectedClaims.Audience[0]) {
			return tok, &claims, ErrAudienceClaimMismatch
		}
	}

	return tok, &claims, nil
}
