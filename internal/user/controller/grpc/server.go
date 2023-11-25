package grpc

import (
	"fmt"
	"net"

	"go.uber.org/zap"

	pb "github.com/agadilkhan/pickup-point-service/pkg/protobuf/userservice/gw"
	"google.golang.org/grpc"
)

type Server struct {
	port       string
	service    *Service
	grpcServer *grpc.Server
	logger     *zap.SugaredLogger
}

func NewServer(
	port string,
	service *Service,
	logger *zap.SugaredLogger,
) *Server {
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	return &Server{
		port:       port,
		service:    service,
		grpcServer: grpcServer,
		logger:     logger,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.port)
	if err != nil {
		return fmt.Errorf("failed to listen grpc port: %s", s.port)
	}

	pb.RegisterUserServiceServer(s.grpcServer, s.service)

	go func() {
		if err := s.grpcServer.Serve(listener); err != nil {
			s.logger.Errorf("failed to serve grpc server err: %v", err)
		}
	}()
	return nil
}

func (s *Server) Close() {
	s.grpcServer.GracefulStop()
}
