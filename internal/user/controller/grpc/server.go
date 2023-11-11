package grpc

import (
	"fmt"
	pb "github.com/agadilkhan/pickup-point-service/pkg/protobuf/userservice/gw"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	port       string
	service    *Service
	grpcServer *grpc.Server
}

func NewServer(
	port string,
	service *Service,
) *Server {
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	return &Server{
		port:       port,
		service:    service,
		grpcServer: grpcServer,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.port)
	if err != nil {
		return fmt.Errorf("failed to listen grpc port: %s", s.port)
	}

	pb.RegisterUserServiceServer(s.grpcServer, s.service)

	go s.grpcServer.Serve(listener)

	return nil
}

func (s *Server) Close() {
	s.grpcServer.GracefulStop()
}
