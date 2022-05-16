package main

import (
	"context"
	"github.com/byorty/enterprise-application/application/order"
	"github.com/byorty/enterprise-application/common/adapters/logger"
	"github.com/byorty/enterprise-application/common/adapters/server/grpc"
	pbv1 "github.com/byorty/enterprise-application/gen/proto/v1"
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
	err := server.Register(grpc.Descriptor{
		Server:               order.NewOrderServiceServer(),
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
