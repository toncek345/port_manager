package port

import (
	"context"
	"errors"
	"testing"

	pb "github.com/toncek345/port_manager/internal/portdomain/proto"
	"github.com/toncek345/port_manager/internal/portdomain/service/port"
	"github.com/stretchr/testify/assert"
)

func Float64ToPtr(f float64) *float64 {
	return &f
}

func TestGetPort(t *testing.T) {
	tests := []struct {
		testName    string
		serviceImpl port.Service
		inObj *pb.GetPortRequest

		expectedOut *pb.Port
		expectedErr bool
	}{
		{
			testName: "service error",
			serviceImpl: &port.Mock{
				GetPortFn: func(context.Context, int64) (*port.Port, error){
					return nil, errors.New("some err")
				},
			},
			inObj: &pb.GetPortRequest{
				PortId: 33,
			},

			expectedOut: nil,
			expectedErr: true,
		},
		{
			testName: "ok",
			serviceImpl: &port.Mock{
				GetPortFn: func(_ context.Context, id int64) (*port.Port, error){
					return &port.Port{
						ID: id,
						IDStr: "idstr",
						Name: "name",
						City: "city",
						Country: "country",
						CoordinatesLat: Float64ToPtr(1.2),
						CoordinatesLon: Float64ToPtr(2.1),
						Provice: "province",
						Timezone: "timezone",
						Code: "code",
						Regions: "regions",
						Unlocs: "unlocs",
						Alias: "alias",
					},nil
				},
			},
			inObj: &pb.GetPortRequest{
				PortId: 33,
			},

			expectedErr: false,
			expectedOut: &pb.Port{
				Id: 33,
				IdStr: "idstr",
				Name: "name",
				City: "city",
				Country: "country",
				Coordinates: []float64{1.2, 2.1},
				Provice: "province",
				Timezone: "timezone",
				Code: "code",
				Regions: []string{"regions"},
				Unlocs: []string{"unlocs"},
				Alias: []string{"alias"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T){
			s := New(test.serviceImpl)

			out, err := s.GetPort(context.Background(), test.inObj)
			if test.expectedErr {
				assert.True(t, true, err!=nil)
			}
			assert.Equal(t, test.expectedOut, out)
		})
	}
}
