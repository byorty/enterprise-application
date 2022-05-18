package main

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/log"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	productapp "github.com/byorty/enterprise-application/pkg/product/infra/app"
	productsrcimpl "github.com/byorty/enterprise-application/pkg/product/infra/service"
)

func main() {
	l := log.NewDefaultLogger()
	server := grpc.NewServer(
		context.Background(),
		l,
		grpc.Config{
			HttpPort: 8080,
			GrpcPort: 8181,
		},
	)

	productService := productsrcimpl.NewProductService()
	err := server.Register(grpc.Descriptor{
		Server:               productapp.NewServer(productService),
		GRPCRegistrar:        pbv1.RegisterProductServiceServer,
		GRPCGatewayRegistrar: pbv1.RegisterProductServiceHandlerFromEndpoint,
	})
	if err != nil {
		l.Fatal(err)
	}

	err = server.Start()
	if err != nil {
		l.Fatal(err)
	}
}
