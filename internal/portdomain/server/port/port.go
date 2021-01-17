package port

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/toncek345/port_manager/internal/portdomain/proto"
	portService "github.com/toncek345/port_manager/internal/portdomain/service/port"
)

type Server struct {
	portService portService.Service
	pb.UnimplementedPortDomainServer
}

func New(portService portService.Service) *Server {
	if portService == nil {
		log.Fatalln("port service is required")
	}

	return &Server{
		portService: portService,
	}
}

func (s *Server) GetPort(ctx context.Context, in *pb.GetPortRequest) (*pb.Port, error) {
	port, err := s.portService.GetPort(ctx, in.PortId)
	if err != nil {
		return nil, fmt.Errorf("finding port: %w", err)
	}

	return &pb.Port{
		Id:      port.ID,
		IdStr:   port.IDStr,
		Name:    port.Name,
		City:    port.City,
		Country: port.Country,
		Coordinates: func() []float64 {
			if port.CoordinatesLat == nil || port.CoordinatesLon == nil {
				return []float64{}
			}

			return []float64{*port.CoordinatesLat, *port.CoordinatesLon}
		}(),
		Provice:  port.Provice,
		Timezone: port.Timezone,
		Code:     port.Code,
		Regions:  strings.Split(port.Regions, ","),
		Unlocs:   strings.Split(port.Unlocs, ","),
		Alias:    strings.Split(port.Alias, ","),
	}, nil
}

func (s *Server) Upsert(in pb.PortDomain_UpsertServer) error {
	for {
		port, err := in.Recv()
		if err != nil {
			if err == io.EOF {
				return in.SendAndClose(&empty.Empty{})
			}

			return fmt.Errorf("receive err: %w", err)
		}

		if err := s.portService.UpsertPort(
			in.Context(),
			&portService.Port{
				IDStr:   port.IdStr,
				Name:    port.Name,
				City:    port.City,
				Country: port.Country,

				CoordinatesLat: func() *float64 {
					if len(port.Coordinates) != 2 {
						return nil
					}
					return &port.Coordinates[1]
				}(),
				CoordinatesLon: func() *float64 {
					if len(port.Coordinates) != 2 {
						return nil
					}
					return &port.Coordinates[0]
				}(),

				Provice:  port.Provice,
				Timezone: port.Timezone,
				Code:     port.Code,
				Regions:  strings.Join(port.Regions, ","),
				Unlocs:   strings.Join(port.Unlocs, ","),
				Alias:    strings.Join(port.Alias, ","),
			},
		); err != nil {
			log.Printf("upserting port: %s", err.Error())
			return fmt.Errorf("upserting port: %w", err)
		}
	}
}
