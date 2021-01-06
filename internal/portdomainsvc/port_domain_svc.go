package portdomainsvc

import (
	"fmt"
	"io"
	"net"
	"strings"

	"log"

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

type PortServer struct {
	portService services.PortService
	pb.UnimplementedPortDomainServer
}

func (s *PortServer) Upsert(in pb.PortDomain_UpsertServer) error {
	for {
		port, err := in.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}

			return fmt.Errorf("receive err: %w", err)
		}

		if err := s.portService.UpsertPort(
			in.Context(),
			&services.Port{
				IDStr:   port.IdStr,
				Name:    port.Name,
				City:    port.City,
				Country: port.Country,

				CoordinatesLat: func() *float64 {
					if len(port.Coordinates) != 2 {
						return nil
					}
					return &port.Coordinates[0]
				}(),
				CoordinatesLon: func() *float64 {
					if len(port.Coordinates) != 2 {
						return nil
					}
					return &port.Coordinates[1]
				}(),

				Provice:  port.Provice,
				Timezone: port.Timezone,
				Code:     port.Code,
				Regions:  strings.Join(port.Regions, ","),
				Unlocs:   strings.Join(port.Unlocs, ","),
				Alias:    strings.Join(port.Alias, ","),
			},
		); err != nil {
			log.Printf("upserting port: %s", err)
			return fmt.Errorf("upserting port: %w", err)
		}
	}
}

func New(port int, db *sqlx.DB) *Service {
	return &Service{
		port:   port,
		db:     db,
		server: grpc.NewServer(),
	}
}
