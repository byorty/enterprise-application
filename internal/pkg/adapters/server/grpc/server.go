package grpc

import (
	"context"
	"fmt"
	"github.com/byorty/enterprise-application/internal/pkg/adapters/logger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"reflect"
)

type Server interface {
	Register(descriptor Descriptor) error
	Start() error
}

type Descriptor struct {
	Server               interface{}
	GRPCRegistrar        interface{}
	GRPCGatewayRegistrar func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error
}

func NewServer(
	ctx context.Context,
	logger logger.Logger,
	cfg Config,
) Server {
	srv := &server{
		ctx:        ctx,
		cfg:        cfg,
		logger:     logger.Named("grpc"),
		grpcServer: grpc.NewServer(),
		mux:        runtime.NewServeMux(),
		opts: []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
		errors: make(chan error, 1),
	}

	return srv
}

type server struct {
	ctx           context.Context
	logger        logger.Logger
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
	go func(logger logger.Logger) {
		netAddress := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.GrpcPort)

		logger.Infof("start server at %s", netAddress)
		socket, err := net.Listen("tcp", netAddress)
		if err != nil {
			s.errors <- err
			return
		}

		s.errors <- s.grpcServer.Serve(socket)
	}(s.logger.Named("grpc_server"))

	go func(logger logger.Logger) {
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
