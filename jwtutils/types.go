package jwtutils

import "github.com/cristalhq/jwt/v4"

type Jwt struct {
	OK    bool   `json:"ok"`
	Error string `json:"error"`
	Jwt   string `json:"jwt"`
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
