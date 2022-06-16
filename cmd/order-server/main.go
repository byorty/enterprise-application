package main

import (
	commonadap "github.com/byorty/enterprise-application/pkg/common/adapter"
	"github.com/byorty/enterprise-application/pkg/common/adapter/application"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	orderadap "github.com/byorty/enterprise-application/pkg/order/infra"
	orderapp "github.com/byorty/enterprise-application/pkg/order/infra/app"
	productsadap "github.com/byorty/enterprise-application/pkg/product/infra"
)

func main() {
	app := application.New(
		commonadap.Constructors,
		productsadap.Constructors,
		orderadap.Constructors,
		orderapp.NewFxOrderServiceServer,
	)
	app.Demonize(func(
		server grpc.Server,
		descriptor grpc.Descriptor,
	) error {
		err := server.Register(descriptor)
		if err != nil {
			return err
		}

		err = server.Start()
		if err != nil {
			return err
		}

		return nil
	})
}
