package auth

import (
	"context"
	"fmt"
	"github.com/byorty/enterprise-application/pkg/common/adapter/application"
	"github.com/byorty/enterprise-application/pkg/common/adapter/log"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	"github.com/casbin/casbin/v2"
)

type Enforcer interface {
	Enforce(ctx context.Context, args ...interface{}) (bool, error)
	HasPolicy(ctx context.Context, params ...interface{}) bool
}

type enforcer struct {
	enforcer *casbin.Enforcer
	logger   log.Logger
}

func NewFxEnforcer(
	logger log.Logger,
	configProvider application.Provider,
) (Enforcer, error) {
	var cfg Config
	err := configProvider.PopulateByKey("enforcer", &cfg)
	if err != nil {
		return nil, err
	}

	casbinEnforcer, err := casbin.NewEnforcer(cfg.ModelFile, cfg.PolicyFile)
	if err != nil {
		return nil, err
	}

	return &enforcer{
		enforcer: casbinEnforcer,
		logger:   logger.Named("enforcer"),
	}, nil
}

func (e *enforcer) Enforce(ctx context.Context, args ...interface{}) (bool, error) {
	casted, err := e.transform(args...)
	if err != nil {
		return false, err
	}
	return e.enforcer.Enforce(casted...)
}

func (e *enforcer) HasPolicy(ctx context.Context, params ...interface{}) bool {
	casted, err := e.transform(params...)
	if err != nil {
		return false
	}
	return e.enforcer.HasPolicy(casted...)
}

func (e *enforcer) transform(args ...interface{}) ([]interface{}, error) {
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
