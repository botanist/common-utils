package jwtutils

import (
	"crypto/subtle"
	"time"

	"github.com/cristalhq/jwt/v4"
)

type Jwt struct {
	OK    bool   `json:"ok"`
	Error string `json:"error"`
	Jwt   string `json:"jwt"`
}

type Claims interface {
	IsIssuer(string) bool
	IsSubject(string) bool
	IsValidAt(time.Time) bool
	IsForAudience(string) bool
	IsValidExpiresAt(time.Time) bool
	IsValidIssuedAt(time.Time) bool
	IsValidNotBefore(time.Time) bool
}

type EntityClaims struct {
	// ID claim provides a unique identifier for the JWT.
	ID        string           `json:"jti,omitempty"`
	Audience  jwt.Audience     `json:"aud,omitempty"`
	Issuer    string           `json:"iss,omitempty"`
	Subject   string           `json:"sub,omitempty"`
	ExpiresAt *jwt.NumericDate `json:"exp,omitempty"`
	IssuedAt  *jwt.NumericDate `json:"iat,omitempty"`
	NotBefore *jwt.NumericDate `json:"nbf,omitempty"`

	Shard string `json:"sid,omitempty"`
}

// Based on github.com/cristalhq/jwt/claims.go

// IsForAudience reports whether token has a given audience.
func (sc *EntityClaims) IsForAudience(audience string) bool {
	for _, aud := range sc.Audience {
		if constTimeEqual(aud, audience) {
			return true
		}
	}
	return false
}

// IsIssuer reports whether token has a given issuer.
func (sc *EntityClaims) IsIssuer(issuer string) bool {
	return constTimeEqual(sc.Issuer, issuer)
}

// IsSubject reports whether token has a given subject.
func (sc *EntityClaims) IsSubject(subject string) bool {
	return constTimeEqual(sc.Subject, subject)
}

// IsValidExpiresAt reports whether a token isn't expired at a given time.
func (sc *EntityClaims) IsValidExpiresAt(now time.Time) bool {
	return sc.ExpiresAt == nil || sc.ExpiresAt.After(now)
}

// IsValidNotBefore reports whether a token isn't used before a given time.
func (sc *EntityClaims) IsValidNotBefore(now time.Time) bool {
	return sc.NotBefore == nil || sc.NotBefore.Before(now)
}

// IsValidIssuedAt reports whether a token was created before a given time.
func (sc *EntityClaims) IsValidIssuedAt(now time.Time) bool {
	return sc.IssuedAt == nil || sc.IssuedAt.Before(now)
}

// IsValidAt reports whether a token is valid at a given time.
func (sc *EntityClaims) IsValidAt(now time.Time) bool {
	return sc.IsValidExpiresAt(now) && sc.IsValidNotBefore(now) && sc.IsValidIssuedAt(now)
}

func constTimeEqual(a, b string) bool {
	return subtle.ConstantTimeCompare([]byte(a), []byte(b)) == 1
}
