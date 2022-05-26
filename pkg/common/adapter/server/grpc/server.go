package grpc

import (
	"context"
	"fmt"
	"github.com/byorty/enterprise-application/pkg/common/adapter/log"
	httpruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	graphqlruntime "github.com/ysugimoto/grpc-graphql-gateway/runtime"
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
	Server                  interface{}
	GRPCRegistrar           interface{}
	GRPCGatewayRegistrar    func(context.Context, *httpruntime.ServeMux, string, []grpc.DialOption) (err error)
	GraphqlGatewayRegistrar func(*graphqlruntime.ServeMux) (err error)
}

func NewServer(
	ctx context.Context,
	logger log.Logger,
	cfg Config,
) Server {
	srv := &server{
		ctx:        ctx,
		cfg:        cfg,
		logger:     logger.Named("grpc"),
		grpcServer: grpc.NewServer(),
		httpMux:    httpruntime.NewServeMux(),
		graphqlMux: graphqlruntime.NewServeMux(),
		opts: []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
		errors: make(chan error, 1),
	}

	return srv
}

type server struct {
	ctx           context.Context
	logger        log.Logger
	cfg           Config
	grpcServer    *grpc.Server
	gatewayServer *http.Server
	httpMux       *httpruntime.ServeMux
	graphqlMux    *graphqlruntime.ServeMux
	opts          []grpc.DialOption
	errors        chan error
}

func (s *server) Register(descriptor Descriptor) error {
	if descriptor.GRPCRegistrar != nil {
		reflect.ValueOf(descriptor.GRPCRegistrar).Call([]reflect.Value{
			reflect.ValueOf(s.grpcServer),
			reflect.ValueOf(descriptor.Server),
		})
	}

	if descriptor.GRPCGatewayRegistrar != nil {
		err := descriptor.GRPCGatewayRegistrar(s.ctx, s.httpMux, fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.GrpcPort), s.opts)
		if err != nil {
			return err
		}
	}

	if descriptor.GraphqlGatewayRegistrar != nil {
		err := descriptor.GraphqlGatewayRegistrar(s.graphqlMux)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *server) Start() error {
	if s.cfg.GrpcPort > 0 {
		go func(logger log.Logger) {
			netAddress := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.GrpcPort)

			logger.Infof("start grpc server at %s", netAddress)
			socket, err := net.Listen("tcp", netAddress)
			if err != nil {
				s.errors <- err
				return
			}

			s.errors <- s.grpcServer.Serve(socket)
		}(s.logger.Named("grpc_server"))
	}

	if s.cfg.HttpPort > 0 {
		go func(logger log.Logger) {
			netAddress := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.HttpPort)

			httpServer := &http.Server{
				Addr:    netAddress,
				Handler: s.httpMux,
			}

			logger.Infof("start http server at %s", netAddress)

			s.errors <- httpServer.ListenAndServe()
		}(s.logger.Named("http_server"))
	}

	if s.cfg.GraphqlPort > 0 {
		go func(logger log.Logger) {
			netAddress := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.GraphqlPort)

			graphqlServer := &http.Server{
				Addr:    netAddress,
				Handler: s.graphqlMux,
			}

			logger.Infof("start graphql server at %s", netAddress)

			s.errors <- graphqlServer.ListenAndServe()
		}(s.logger.Named("graphql_server"))
	}

	return <-s.errors
}
