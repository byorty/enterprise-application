package main

import (
	"context"
	"github.com/byorty/enterprise-application/internal/pkg/adapters/logger"
	"github.com/byorty/enterprise-application/internal/pkg/adapters/server/grpc"
	pbv1 "github.com/byorty/enterprise-application/internal/pkg/gen/api/proto/v1"
	productapp "github.com/byorty/enterprise-application/internal/product/infra/app"
	productsrcimpl "github.com/byorty/enterprise-application/internal/product/infra/service"
)

func main() {
	l := logger.NewDefaultLogger()
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
