package jwtutil_test

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"github.com/byorty/enterprise-application/pkg/common/adapter/application"
	"github.com/byorty/enterprise-application/pkg/common/adapter/jwtutil"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"go.uber.org/config"
	"os"
	"strings"
	"testing"
	"time"
)

func TestJwtHelperSuite(t *testing.T) {
	suite.Run(t, new(JwtHelperSuite))
}

type JwtHelperSuite struct {
	suite.Suite
	helper jwtutil.Helper
}

func (s *JwtHelperSuite) SetupSuite() {
	dir, err := os.Getwd()
	s.Nil(err)

	reader := strings.NewReader(fmt.Sprintf(`
ssl:
  private_key_file: %s/private.key.pem
  public_key_file: %s/public.key.pem
`, dir, dir))

	provider, err := application.NewProviderByOptions(config.Source(reader))
	s.Nil(err)

	s.helper, err = jwtutil.NewFxHelper(provider)
	s.Nil(err)
}

func (s *JwtHelperSuite) TestAll() {
	claims := &jwtutil.SessionClaims{}
	err := s.helper.Parse(randomdata.Alphanumeric(5), claims)
	s.Contains(err.Error(), "token contains an invalid number of segments")

	claims = &jwtutil.SessionClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(-1 * time.Hour).Unix(),
		},
		Session: pbv1.Session{
			Uuid:  "",
			Group: -1,
		},
	}
	token, err := s.helper.CreateToken(claims)
	s.NotEmpty(token)
	s.Nil(err)

	err = s.helper.Parse(token, claims)
	s.Contains(err.Error(), "token is expired by")

	claims.StandardClaims.ExpiresAt = time.Now().Add(time.Hour).Unix()
	token, err = s.helper.CreateToken(claims)
	s.NotEmpty(token)
	s.Nil(err)

	err = s.helper.Parse(token, claims)
	s.Contains(err.Error(), "uuid is invalid")

	claims.Uuid = uuid.NewString()
	token, err = s.helper.CreateToken(claims)
	s.NotEmpty(token)
	s.Nil(err)

	err = s.helper.Parse(token, claims)
	s.Contains(err.Error(), "group is invalid")

	claims.Group = pbv1.UserGroupCustomer
	token, err = s.helper.CreateToken(claims)
	s.NotEmpty(token)
	s.Nil(err)

	err = s.helper.Parse(token, claims)
	s.Nil(err)
}
