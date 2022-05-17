package main

import (
	"context"
	"github.com/byorty/enterprise-application/internal/pkg/adapters/logger"
	"github.com/byorty/enterprise-application/internal/pkg/adapters/server/grpc"
	pbv1 "github.com/byorty/enterprise-application/internal/pkg/gen/api/proto/v1"
	userapp "github.com/byorty/enterprise-application/internal/user/infra/app"
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
		Server:               userapp.NewServer(),
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
