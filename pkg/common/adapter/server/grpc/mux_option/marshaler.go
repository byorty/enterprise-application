package mux_option

import (
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/encoding/protojson"
)

func NewFxMarshalerOption() grpc.MiddlewareOut {
	return grpc.MiddlewareOut{
		MuxMiddleware: grpc.Middleware{
			MuxOption: runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:   true,
					EmitUnpopulated: true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{},
			}),
		},
	}
}
