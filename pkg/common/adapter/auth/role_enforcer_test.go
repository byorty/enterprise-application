package auth_test

import (
	"fmt"
	"github.com/byorty/enterprise-application/pkg/common/adapter/application"
	"github.com/byorty/enterprise-application/pkg/common/adapter/auth"
	"github.com/byorty/enterprise-application/pkg/common/adapter/log"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	"github.com/stretchr/testify/suite"
	"go.uber.org/config"
	"os"
	"strings"
	"testing"
)

func TestRoleEnforcerSuite(t *testing.T) {
	suite.Run(t, new(RoleEnforcerSuite))
}

type RoleEnforcerSuite struct {
	suite.Suite
	enforcer auth.RoleEnforcer
}

func (s *RoleEnforcerSuite) SetupSuite() {
	dir, err := os.Getwd()
	s.Nil(err)

	reader := strings.NewReader(fmt.Sprintf(`
enforcer:
  model_file: %s/model.conf
  policy_file: %s/policy.csv
`, dir, dir))

	provider, err := application.NewProviderByOptions(config.Source(reader))
	s.Nil(err)

	s.enforcer, err = auth.NewFxRoleEnforcer(
		log.NewFxLogger(),
		provider,
	)
	s.Nil(err)
}

func (s *RoleEnforcerSuite) TestConstructor() {
	reader := strings.NewReader(fmt.Sprintf(`
enforcer:
  model_file: "1234"
  policy_file: "3455"
`))

	provider, err := application.NewProviderByOptions(config.Source(reader))
	s.Nil(err)
	enf, err := auth.NewFxRoleEnforcer(
		log.NewFxLogger(),
		provider,
	)
	s.Nil(enf)
	s.NotNil(err)
}

func (s *RoleEnforcerSuite) TestEnforce() {
	guestSession := pbv1.Session{
		Group: pbv1.UserGroupGuest,
	}

	ok, err := s.enforcer.Enforce(
		guestSession,
		pbv1.RoleUser,
		pbv1.PermissionRead,
	)
	s.Nil(err)
	s.True(ok)

	ok, err = s.enforcer.Enforce(
		guestSession,
		pbv1.RoleOrder,
		pbv1.PermissionRead,
	)
	s.Nil(err)
	s.False(ok)

	customerSession := pbv1.Session{
		Group: pbv1.UserGroupCustomer,
	}

	ok, err = s.enforcer.Enforce(
		customerSession,
		pbv1.RoleUser,
		pbv1.PermissionRead,
	)
	s.Nil(err)
	s.True(ok)

	ok, err = s.enforcer.Enforce(
		customerSession,
		pbv1.RoleOrder,
		pbv1.PermissionRead,
	)
	s.Nil(err)
	s.True(ok)
}
