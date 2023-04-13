package models

import (
	"crypto"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenKeys struct {
	PrivateKey crypto.PrivateKey
	PublicKey  crypto.PublicKey
}

type Token struct {
	AccessToken string
}

type CustomClaims struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

type CustomJWTClaims struct {
	jwt.RegisteredClaims
	CustomClaims
}
