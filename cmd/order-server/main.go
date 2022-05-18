package main

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/log"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	orderapp "github.com/byorty/enterprise-application/pkg/order/infra/app"
	ordersrvimpl "github.com/byorty/enterprise-application/pkg/order/infra/service"
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
