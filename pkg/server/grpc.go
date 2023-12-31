// Package server is package that helps to manage server instance of application
package server

import (
	"context"
	"crypto/tls"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"secretKeeper/internal/server/config"
	"secretKeeper/pkg/cert"
)

type Service interface {
	RegisterService(grpc.ServiceRegistrar)
}

type GrpcServer struct {
	cfg                config.Config
	logger             *zap.Logger
	server             *grpc.Server
	services           []Service
	unaryInterceptors  []grpc.UnaryServerInterceptor
	streamInterceptors []grpc.StreamServerInterceptor
}

// GrpcServerOption - is callback function that applies an option to GrpcServer.
type GrpcServerOption func(*GrpcServer)

// WithServerConfig - adds config.Config to GrpcServer.
func WithServerConfig(c config.Config) GrpcServerOption {
	return func(server *GrpcServer) {
		server.cfg = c
	}
}

// WithLogger - adds *zap.Logger to GrpcServer.
func WithLogger(z *zap.Logger) GrpcServerOption {
	return func(server *GrpcServer) {
		server.logger = z
	}
}

// WithServices - adds []Service to GrpcServer.
func WithServices(s ...Service) GrpcServerOption {
	return func(server *GrpcServer) {
		server.services = s
	}
}

// WithStreamInterceptors - add []grpc.StreamServerInterceptor to GrpcServer.
func WithStreamInterceptors(in ...grpc.StreamServerInterceptor) GrpcServerOption {
	return func(server *GrpcServer) {
		server.streamInterceptors = append(server.streamInterceptors, in...)
	}
}

// WithUnaryInterceptors - add []grpc.UnaryServerInterceptor to GrpcServer.
func WithUnaryInterceptors(in ...grpc.UnaryServerInterceptor) GrpcServerOption {
	return func(server *GrpcServer) {
		server.unaryInterceptors = append(server.unaryInterceptors, in...)
	}
}

// NewGrpcServer - creates new GrpcServer with options via provided GrpcServerOption.
func NewGrpcServer(opts ...GrpcServerOption) *GrpcServer {
	s := &GrpcServer{}

	for _, option := range opts {
		option(s)
	}

	return s
}

// RegisterServices - adds a service to gRPC server via, RegisterService function which is called on each provided
// Service.
func (s *GrpcServer) RegisterServices(services ...Service) {
	for _, service := range services {
		service.RegisterService(s.server)
	}
}

// Start - starts gRPC server with enabled TLS on port from config.Config.
func (s *GrpcServer) Start(cancel context.CancelFunc) {
	sslConf, err := cert.NewSSLConfigService().LoadServerCertificate(s.cfg)

	conn, errListen := tls.Listen("tcp", ":"+s.cfg.Port, sslConf)
	if errListen != nil {
		s.logger.Error(errListen.Error())
	}

	s.server = grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			s.streamInterceptors...,
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			s.unaryInterceptors...,
		)),
	)

	s.RegisterServices(s.services...)

	go func() {
		err = s.server.Serve(conn)
		if err != nil {
			s.logger.Error(err.Error())
			cancel()
		}
	}()

	s.logger.Sugar().Infof("gRPC server is running on %s port", s.cfg.Port)
}

// Stop - gracefully stops gRPC server.
func (s *GrpcServer) Stop() {
	s.logger.Info("Gracefully stopping gRPC server")

	s.server.GracefulStop()

	s.logger.Info("gRPC server stopped")
}
