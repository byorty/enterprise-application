package mux_option

import (
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

func NewFxMarshalerOption() grpc.MiddlewareOut {
	return grpc.MiddlewareOut{
		MuxMiddleware: grpc.Middleware{
			MuxOption: runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
				OrigName: true,
			}),
		},
	}
}
