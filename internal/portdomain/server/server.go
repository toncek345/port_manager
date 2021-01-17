package server

import (
	"net"

	pb "github.com/toncek345/port_manager/internal/portdomain/proto"
	"github.com/toncek345/port_manager/internal/portdomain/server/port"
	"github.com/toncek345/port_manager/internal/portdomain/service"
	"google.golang.org/grpc"
)

type Server struct {
	server *grpc.Server
}

func New(services *service.Services) *Server {
	grpcServer := grpc.NewServer()

	pb.RegisterPortDomainServer(grpcServer, port.New(services.Port))

	return &Server{
		server: grpcServer,
	}
}

func (s *Server) Serve(listener net.Listener) error {
	return s.server.Serve(listener)
}

func (s *Server) Close() {
	s.server.GracefulStop()
}
