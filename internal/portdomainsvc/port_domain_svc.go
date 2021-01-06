package portdomainsvc

import (
	"fmt"
	"net"

	"github.com/jmoiron/sqlx"
	pb "github.com/toncek345/port_manager/internal/portdomainsvc/grpc"
	"google.golang.org/grpc"
)

type Service struct {
	port   int
	db     *sqlx.DB
	server *grpc.Server
}

func (s *Service) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	pb.RegisterPortDomainServer(s.server, &PortServer{})

	if err := s.server.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}

func (s *Service) Stop() {
	s.server.Stop()
}

type PortServer struct {
	pb.UnimplementedPortDomainServer
}

func (s *PortServer) Upsert(in pb.PortDomain_UpsertServer) error {
	fmt.Printf("HEEEEEEEEEEEEELOOOOOOOOOOOOOOOO GRRRPC\n\n\n")

	return nil
}

func (s *PortServer) mustEmbedUnimplementedPortDomainServer() {}

func New(port int, db *sqlx.DB) *Service {
	return &Service{
		port:   port,
		db:     db,
		server: grpc.NewServer(),
	}
}
