package clientapi

import pb "github.com/toncek345/port_manager/internal/portdomain/proto"

type Port struct {
	ID          int64     `json:"id"`
	IDStr       string    `json:"IDStr"`
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Coordinates []float64 `json:"coordinates"`
	Provice     string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Code        string    `json:"code"`
	Regions     []string  `json:"regions"`
	Unlocs      []string  `json:"unlocs"`
	Alias       []string  `json:"alias"`
}

func PortToPortProto(port *Port) *pb.Port {
	return &pb.Port{
		Id:          port.ID,
		IdStr:       port.IDStr,
		Name:        port.Name,
		City:        port.City,
		Country:     port.Country,
		Coordinates: port.Coordinates,
		Provice:     port.Provice,
		Timezone:    port.Timezone,
		Code:        port.Code,
		Regions:     port.Regions,
		Unlocs:      port.Unlocs,
		Alias:       port.Alias,
	}
}

func PortProtoToPort(port *pb.Port) *Port {
	return &Port{
		ID:          port.Id,
		IDStr:       port.IdStr,
		Name:        port.Name,
		City:        port.City,
		Country:     port.Country,
		Coordinates: port.Coordinates,
		Provice:     port.Provice,
		Timezone:    port.Timezone,
		Code:        port.Code,
		Regions:     port.Regions,
		Unlocs:      port.Unlocs,
		Alias:       port.Alias,
	}
}
