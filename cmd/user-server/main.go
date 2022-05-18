package main

import (
	"context"
	"github.com/byorty/enterprise-application/internal/pkg/adapters/logger"
	"github.com/byorty/enterprise-application/internal/pkg/adapters/server/grpc"
	pbv1 "github.com/byorty/enterprise-application/internal/pkg/gen/api/proto/v1"
	productsrcimpl "github.com/byorty/enterprise-application/internal/product/infra/service"
	userapp "github.com/byorty/enterprise-application/internal/user/infra/app"
	usersrvimpl "github.com/byorty/enterprise-application/internal/user/infra/service"
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
	userService := usersrvimpl.NewUserService()
	userProductService := usersrvimpl.NewUserProductService(userService, productService)
	err := server.Register(grpc.Descriptor{
		Server:               userapp.NewServer(userService, userProductService),
		GRPCRegistrar:        pbv1.RegisterUserServiceServer,
		GRPCGatewayRegistrar: pbv1.RegisterUserServiceHandlerFromEndpoint,
	})
	if err != nil {
		l.Fatal(err)
	}

	err = server.Start()
	if err != nil {
		l.Fatal(err)
	}
}
