package main

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/log"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
)

func main() {
	l := log.NewDefaultLogger()
	server := grpc.NewServer(
		context.Background(),
		l,
		grpc.Config{
			GraphqlPort: 8282,
		},
	)

	err := server.Register(grpc.Descriptor{
		GraphqlGatewayRegistrar: pbv1.RegisterProductServiceGraphql,
	})
	if err != nil {
		l.Fatal(err)
	}

	err = server.Register(grpc.Descriptor{
		GraphqlGatewayRegistrar: pbv1.RegisterUserServiceGraphql,
	})
	if err != nil {
		l.Fatal(err)
	}

	err = server.Start()
	if err != nil {
		l.Fatal(err)
	}
}
