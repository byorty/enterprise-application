package main

import (
	commonadap "github.com/byorty/enterprise-application/pkg/common/adapter"
	"github.com/byorty/enterprise-application/pkg/common/adapter/application"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	orderadap "github.com/byorty/enterprise-application/pkg/order/infra"
	productsadap "github.com/byorty/enterprise-application/pkg/product/infra"
	productapp "github.com/byorty/enterprise-application/pkg/product/infra/app"
)

func main() {
	app := application.New(
		commonadap.Constructors,
		productsadap.Constructors,
		orderadap.Constructors,
		productapp.NewFxProductServiceServer,
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
