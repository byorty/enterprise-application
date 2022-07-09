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
	grpc.NewFxMethodDescriptorMap,
	grpc_option.NewFxPanicOption,
	grpc_option.NewFxAuthOption,
	grpc_option.NewFxEnforcerOption,
	grpc_option.NewFxValidatorOption,
	mux_option.NewFxMarshalerOption,
	auth.NewFxHelper,
	auth.NewFxRoleEnforcer,
	//validator.NewFxMap,
	//validator.NewFxMessage,
	//validator.NewFxUUID(
	//	"uuid",
	//	"order_uuid",
	//	"product_uuid",
	//	"user_uuid",
	//),
	//validator.NewFxOptionalUint32(
	//	"offset",
	//	"limit",
	//),
	//validator.NewFxRequiredUint32(
	//	"count",
	//),
)
