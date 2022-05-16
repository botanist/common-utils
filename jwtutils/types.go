package jwtutils

import "github.com/cristalhq/jwt/v4"

type Jwt struct {
	OK    bool   `json:"ok"`
	Error string `json:"error"`
	Jwt   string `json:"jwt"`
}

type EntityClaims struct {
	jwt.RegisteredClaims
	Shard string `json:"sid,omitempty"`
}
