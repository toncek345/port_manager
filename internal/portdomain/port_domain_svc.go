package portdomainsvc

import (
	"fmt"
	"net"

	"github.com/jmoiron/sqlx"
	pb "github.com/toncek345/port_manager/internal/portdomainsvc/grpc"
	"github.com/toncek345/port_manager/internal/portdomainsvc/services"
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

	pb.RegisterPortDomainServer(s.server, &PortServer{portService: &services.PortSQL{DB: s.db}})

	if err := s.server.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}

func (s *Service) Stop() {
	s.server.Stop()
}

func New(port int, db *sqlx.DB) *Service {
	return &Service{
		port:   port,
		db:     db,
		server: grpc.NewServer(),
	}
}
