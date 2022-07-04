package grpc

import (
	"context"
	"fmt"
	"github.com/byorty/enterprise-application/pkg/common/adapter/application"
	"github.com/byorty/enterprise-application/pkg/common/adapter/log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"reflect"
	"sort"
)

type Middleware struct {
	Priority   int
	GrpcOption grpc.UnaryServerInterceptor
	MuxOption  runtime.ServeMuxOption
}

type ByPriority []Middleware

func (b ByPriority) Len() int {
	return len(b)
}

func (b ByPriority) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByPriority) Less(i, j int) bool {
	return b[i].Priority > b[j].Priority
}

type MiddlewareOut struct {
	fx.Out
	GrpcMiddleware Middleware `group:"grpc_middleware"`
	MuxMiddleware  Middleware `group:"mux_middleware"`
}

type Server interface {
	Register(descriptor Descriptor) error
	Start() error
}

type Descriptor struct {
	Server               interface{}
	GRPCRegistrar        interface{}
	GRPCGatewayRegistrar func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error
	MethodDescriptors    []MethodDescriptor
}

type FxServerIn struct {
	fx.In
	Ctx             context.Context
	Logger          log.Logger
	ConfigProvider  application.Provider
	GrpcMiddlewares []Middleware `group:"grpc_middleware"`
	MuxMiddlewares  []Middleware `group:"mux_middleware"`
}

func NewFxServer(
	in FxServerIn,
) (Server, error) {
	var cfg Config
	err := in.ConfigProvider.PopulateByKey("server", &cfg)
	if err != nil {
		return nil, err
	}

	sort.Sort(ByPriority(in.MuxMiddlewares))
	sort.Sort(ByPriority(in.GrpcMiddlewares))

	interceptors := make([]grpc.UnaryServerInterceptor, len(in.GrpcMiddlewares))
	for i, middleware := range in.GrpcMiddlewares {
		interceptors[i] = middleware.GrpcOption
	}

	serverMuxOptions := make([]runtime.ServeMuxOption, len(in.MuxMiddlewares))
	for i, middleware := range in.MuxMiddlewares {
		serverMuxOptions[i] = middleware.MuxOption
	}

	srv := &server{
		ctx:    in.Ctx,
		cfg:    cfg,
		logger: in.Logger.Named("grpc"),
		grpcServer: grpc.NewServer(
			grpc.MaxRecvMsgSize(cfg.MaxReceiveMessageLength),
			grpc.MaxSendMsgSize(cfg.MaxSendMessageLength),
			grpc.ChainUnaryInterceptor(interceptors...),
		),
		mux: runtime.NewServeMux(
			serverMuxOptions...,
		),
		opts: []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultCallOptions(
				grpc.MaxCallRecvMsgSize(cfg.MaxReceiveMessageLength),
				grpc.MaxCallSendMsgSize(cfg.MaxSendMessageLength),
			),
		},
		errors: make(chan error, 1),
	}

	return srv, nil
}

type server struct {
	ctx           context.Context
	logger        log.Logger
	cfg           Config
	grpcServer    *grpc.Server
	gatewayServer *http.Server
	mux           *runtime.ServeMux
	opts          []grpc.DialOption
	errors        chan error
}

func (s *server) Register(descriptor Descriptor) error {
	reflect.ValueOf(descriptor.GRPCRegistrar).Call([]reflect.Value{
		reflect.ValueOf(s.grpcServer),
		reflect.ValueOf(descriptor.Server),
	})
	if descriptor.GRPCGatewayRegistrar != nil {
		return descriptor.GRPCGatewayRegistrar(s.ctx, s.mux, fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.GrpcPort), s.opts)
	}

	return nil
}

func (s *server) Start() error {
	go func(logger log.Logger) {
		netAddress := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.GrpcPort)

		logger.Infof("start server at %s", netAddress)
		socket, err := net.Listen("tcp", netAddress)
		if err != nil {
			s.errors <- err
			return
		}

		s.errors <- s.grpcServer.Serve(socket)
	}(s.logger.Named("grpc_server"))

	go func(logger log.Logger) {
		netAddress := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.HttpPort)

		s.gatewayServer = &http.Server{
			Addr:    netAddress,
			Handler: s.mux,
		}

		logger.Infof("start gateway at %s", netAddress)

		s.errors <- s.gatewayServer.ListenAndServe()
	}(s.logger.Named("http_server"))

	return <-s.errors
}
