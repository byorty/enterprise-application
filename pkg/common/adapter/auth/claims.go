package auth

import (
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type SessionClaims struct {
	jwt.StandardClaims
	pbv1.Session
}

func (c SessionClaims) Valid() error {
	err := c.StandardClaims.Valid()
	if err != nil {
		return err
	}

	if len(c.Uuid) == 0 {
		return errors.New("uuid is invalid")
	}

	if _, ok := pbv1.UserGroup_name[int32(c.Group)]; !ok {
		return errors.New("group is invalid")
	}

	return nil
}
