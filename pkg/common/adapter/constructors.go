package commonadap

import (
	"github.com/byorty/enterprise-application/pkg/common/adapter/application"
	"github.com/byorty/enterprise-application/pkg/common/adapter/log"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	"go.uber.org/fx"
)

var Constructors = fx.Provide(
	application.NewFxProvider,
	log.NewFxLogger,
	grpc.NewFxServer,
)
