package auth

import (
	"fmt"
	"github.com/byorty/enterprise-application/pkg/common/adapter/application"
	"github.com/byorty/enterprise-application/pkg/common/adapter/log"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	"github.com/casbin/casbin/v2"
)

type RoleEnforcer interface {
	Enforce(args ...interface{}) (bool, error)
	HasPolicy(args ...interface{}) bool
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

func (e *roleEnforcer) Enforce(args ...interface{}) (bool, error) {
	casted, err := e.transform(args...)
	if err != nil {
		return false, err
	}
	return e.enforcer.Enforce(casted...)
}

func (e *roleEnforcer) HasPolicy(args ...interface{}) bool {
	casted, err := e.transform(args...)
	if err != nil {
		return false
	}
	return e.enforcer.HasPolicy(casted...)
}

func (e *roleEnforcer) transform(args ...interface{}) ([]interface{}, error) {
	var session pbv1.Session
	var obj pbv1.PermissionObject
	var operation pbv1.PermissionOperation
	var ok bool
	for i, arg := range args {
		switch i {
		case 0:
			session, ok = arg.(pbv1.Session)
			if !ok {
				return nil, fmt.Errorf("%d arg type %T is unexpected ", i, arg)
			}
		case 1:
			obj, ok = arg.(pbv1.PermissionObject)
			if !ok {
				return nil, fmt.Errorf("%d arg type %T is unexpected ", i, arg)
			}
		case 2:
			operation, ok = arg.(pbv1.PermissionOperation)
			if !ok {
				return nil, fmt.Errorf("%d arg type %T is unexpected ", i, arg)
			}
		}
	}

	return []interface{}{
		pbv1.UserGroup_name[int32(session.Group)],
		pbv1.PermissionObject_name[int32(obj)],
		pbv1.PermissionOperation_name[int32(operation)],
	}, nil
}
