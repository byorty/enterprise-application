package jwtutil

import (
	"crypto/rsa"
	"github.com/byorty/enterprise-application/pkg/common/adapter/application"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
)

type Claims interface {
	jwt.Claims
}

type Helper interface {
	Parse(token string, claims Claims) error
	CreateToken(claims Claims) (string, error)
}

func NewFxHelper(
	provider application.Provider,
) (Helper, error) {
	var cfg Config
	err := provider.PopulateByKey("ssl", &cfg)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadFile(cfg.PrivateKeyFile)
	if err != nil {
		return nil, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(buf)
	if err != nil {
		return nil, err
	}

	buf, err = ioutil.ReadFile(cfg.PublicKeyFile)
	if err != nil {
		return nil, err
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(buf)
	if err != nil {
		return nil, err
	}

	return &helper{
		publicKey:  publicKey,
		privateKey: privateKey,
	}, nil
}

type helper struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

func (h *helper) Parse(token string, claims Claims) error {
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return h.publicKey, nil
	})
	if err != nil {
		return err
	}

	return claims.Valid()
}

func (h *helper) CreateToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(h.privateKey)
}
