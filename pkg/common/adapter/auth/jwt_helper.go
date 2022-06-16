package auth

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
)

type Claims interface {
	jwt.Claims
}

type JWTHelper interface {
	Parse(token string, claims Claims) error
	CreateToken(claims Claims) (string, error)
}

func NewFxJWTHelper(publicKey *rsa.PublicKey, privateKey *rsa.PrivateKey) JWTHelper {
	return &jwtHelper{
		publicKey:  publicKey,
		privateKey: privateKey,
	}
}

type jwtHelper struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

func (h *jwtHelper) Parse(token string, claims Claims) error {
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return h.publicKey, nil
	})
	if err != nil {
		return err
	}
	err = claims.Valid()
	if err != nil {
		return err
	}
	return nil
}

func (h *jwtHelper) CreateToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(h.privateKey)
}
