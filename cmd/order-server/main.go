package main

import (
	"context"
	orderapp "github.com/byorty/enterprise-application/internal/order/infra/app"
	ordersrvimpl "github.com/byorty/enterprise-application/internal/order/infra/service"
	"github.com/byorty/enterprise-application/internal/pkg/adapters/logger"
	"github.com/byorty/enterprise-application/internal/pkg/adapters/server/grpc"
	pbv1 "github.com/byorty/enterprise-application/internal/pkg/gen/api/proto/v1"
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
	orderService := ordersrvimpl.NewOrderService(productService)
	err := server.Register(grpc.Descriptor{
		Server:               orderapp.NewOrderServiceServer(orderService),
		GRPCRegistrar:        pbv1.RegisterOrderServiceServer,
		GRPCGatewayRegistrar: pbv1.RegisterOrderServiceHandlerFromEndpoint,
	})
	if err != nil {
		l.Fatal(err)
	}

	err = server.Start()
	if err != nil {
		l.Fatal(err)
	}
}
