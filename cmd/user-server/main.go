package main

import (
	commonadap "github.com/byorty/enterprise-application/pkg/common/adapter"
	"github.com/byorty/enterprise-application/pkg/common/adapter/application"
	"github.com/byorty/enterprise-application/pkg/common/adapter/server/grpc"
	orderadap "github.com/byorty/enterprise-application/pkg/order/infra"
	productsadap "github.com/byorty/enterprise-application/pkg/product/infra"
	useradap "github.com/byorty/enterprise-application/pkg/user/infra"
	userapp "github.com/byorty/enterprise-application/pkg/user/infra/app"
)

func main() {
	app := application.New(
		commonadap.Constructors,
		productsadap.Constructors,
		useradap.Constructors,
		orderadap.Constructors,
		userapp.NewFxUserServiceServer,
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
