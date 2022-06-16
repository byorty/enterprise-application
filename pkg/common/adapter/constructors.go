package commonadap

import (
	"github.com/byorty/enterprise-application/pkg/common/adapter/application"
	"github.com/byorty/enterprise-application/pkg/common/adapter/auth"
	"github.com/byorty/enterprise-application/pkg/common/adapter/log"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc/grpc_option"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc/mux_option"
	"go.uber.org/fx"
)

var Constructors = fx.Provide(
	application.NewFxProvider,
	log.NewFxLogger,
	grpc.NewFxServer,
	grpc_option.NewFxEnforcerOption,
	mux_option.NewFxMarshalerOption,
	auth.NewFxJWTHelper,
	auth.NewFxEnforcer,
)
