package portdomainsvc

import (
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
