package auth

import (
	"errors"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	"github.com/dgrijalva/jwt-go"
)

type SessionClaims struct {
	jwt.StandardClaims
	pbv1.Session
}

func (s SessionClaims) Valid() error {
	err := s.StandardClaims.Valid()
	if err != nil {
		return err
	}

	if len(s.Uuid) == 0 {
		return errors.New("uuid is invalid")
	}

	if _, ok := pbv1.UserGroup_name[int32(s.Group)]; !ok {
		return errors.New("group is invalid")
	}

	return nil
}
