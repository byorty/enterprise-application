package application

import (
	"context"
	"github.com/byorty/enterprise-application/pkg/common/adapter/log"
	"github.com/jessevdk/go-flags"
	"go.uber.org/fx"
	"os"
	"os/signal"
	"syscall"
)

type Application struct {
	ctx     context.Context
	cancel  context.CancelFunc
	fxApp   *fx.App
	logger  log.Logger
	options []fx.Option
}

func New(providers ...interface{}) *Application {
	ctx, cancel := context.WithCancel(context.Background())

	app := &Application{
		options: []fx.Option{
			fx.NopLogger,
		},
		ctx:    ctx,
		cancel: cancel,
	}

	for _, provider := range providers {
		switch p := provider.(type) {
		case fx.Option:
			app.options = append(app.options, p)
		default:
			app.options = append(app.options, fx.Provide(p))
		}
	}

	return app
}

func (a *Application) Run(invoker interface{}) {
	var args Arguments
	_, err := flags.Parse(&args)
	if err != nil {
		panic(err)
	}

	a.options = append(
		a.options,
		fx.Provide(func() context.Context {
			return a.ctx
		}),
		fx.Provide(func() Arguments {
			return args
		}),
		fx.Invoke(invoker),
		fx.Populate(&a.logger),
	)
	a.fxApp = fx.New(a.options...)

	go a.listenSignals()

	startCtx, cancel := context.WithTimeout(a.ctx, fx.DefaultTimeout)
	defer cancel()

	if err = a.fxApp.Start(startCtx); err != nil {
		if a.logger == nil {
			panic(err)
		}

		a.logger.Fatal(err)
	}
}

func (a *Application) Demonize(invoker interface{}) {
	a.Run(invoker)
	<-a.ctx.Done()
}

func (a *Application) Stop() {
	stopCtx, cancel := context.WithTimeout(a.ctx, fx.DefaultTimeout)
	defer cancel()
	err := a.fxApp.Stop(stopCtx)
	if err != nil {
		a.logger.Fatal(err)
	}
	a.fxApp = nil
}

func (a *Application) listenSignals() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for sig := range signals {
		a.logger.Infof("income signal %s", sig)
		a.Stop()
		a.cancel()
		return
	}
}
