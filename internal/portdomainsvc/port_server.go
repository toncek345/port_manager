package portdomainsvc

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	pb "github.com/toncek345/port_manager/internal/portdomainsvc/grpc"
	"github.com/toncek345/port_manager/internal/portdomainsvc/services"
)

type PortServer struct {
	portService services.PortService
	pb.UnimplementedPortDomainServer
}

func (s *PortServer) GetPort(ctx context.Context, in *pb.GetPortRequest) (*pb.Port, error) {
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

			return []float64{*port.CoordinatesLon, *port.CoordinatesLat}
		}(),
		Provice:  port.Provice,
		Timezone: port.Timezone,
		Code:     port.Code,
		Regions:  strings.Split(port.Regions, ","),
		Unlocs:   strings.Split(port.Unlocs, ","),
		Alias:    strings.Split(port.Alias, ","),
	}, nil
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
