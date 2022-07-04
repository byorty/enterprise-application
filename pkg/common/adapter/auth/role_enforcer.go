package auth

import (
	"github.com/byorty/enterprise-application/pkg/common/adapter/application"
	"github.com/byorty/enterprise-application/pkg/common/adapter/log"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	"github.com/casbin/casbin/v2"
	"strings"
)

type RoleEnforcer interface {
	Enforce(session pbv1.Session, role pbv1.Role, permission pbv1.Permission) (bool, error)
}

type roleEnforcer struct {
	enforcer *casbin.Enforcer
	logger   log.Logger
}

func NewFxRoleEnforcer(
	logger log.Logger,
	configProvider application.Provider,
) (RoleEnforcer, error) {
	var cfg Config
	err := configProvider.PopulateByKey("enforcer", &cfg)
	if err != nil {
		return nil, err
	}

	casbinEnforcer, err := casbin.NewEnforcer(cfg.ModelFile, cfg.PolicyFile)
	if err != nil {
		return nil, err
	}

	return &roleEnforcer{
		enforcer: casbinEnforcer,
		logger:   logger.Named("role_enforcer"),
	}, nil
}

func (e *roleEnforcer) Enforce(session pbv1.Session, role pbv1.Role, permission pbv1.Permission) (bool, error) {
	return e.enforcer.Enforce(
		strings.ReplaceAll(pbv1.UserGroup_name[int32(session.Group)], "USER_GROUP_", ""),
		pbv1.Role_name[int32(role)],
		pbv1.Permission_name[int32(permission)],
	)
}
