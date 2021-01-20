package clientapi

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	pb "github.com/toncek345/port_manager/internal/portdomain/proto"
	"google.golang.org/grpc"
)

func TestGetPort(t *testing.T) {
	t.Run("get port works", func(t *testing.T) {
		serviceMock := &PortServiceMock{
			GetPortFn: func(_ context.Context, in *pb.GetPortRequest, _ ...grpc.CallOption) (*pb.Port, error) {
				assert.Equal(t, int64(55), in.PortId)
				return &pb.Port{
					Id:          55,
					IdStr:       "55",
					Name:        "name",
					City:        "city",
					Country:     "country",
					Coordinates: []float64{1, 2},
					Provice:     "province",
					Timezone:    "timezone",
					Code:        "code",
					Regions:     []string{"region1", "region2"},
					Unlocs:      []string{"unlock1", "unlock2"},
					Alias:       []string{"alias1", "alias2"},
				}, nil
			},
		}
		service := New(serviceMock)

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/ports?id=55", nil)
		service.Router.ServeHTTP(rec, req)

		resp := rec.Result()
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		data, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err)

		var port Port
		err = json.Unmarshal(data, &port)
		assert.Nil(t, err)

		assert.Equal(
			t,
			&Port{
				ID:          55,
				IDStr:       "55",
				Name:        "name",
				City:        "city",
				Country:     "country",
				Coordinates: []float64{1, 2},
				Provice:     "province",
				Timezone:    "timezone",
				Code:        "code",
				Regions:     []string{"region1", "region2"},
				Unlocs:      []string{"unlock1", "unlock2"},
				Alias:       []string{"alias1", "alias2"},
			},
			&port)
	})

	t.Run("bad id", func(t *testing.T) {
		service := New(&PortServiceMock{})
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/ports?id=asdf", nil)
		service.Router.ServeHTTP(rec, req)

		resp := rec.Result()
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		data, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err)

		var errResp ErrorResponse
		err = json.Unmarshal(data, &errResp)
		assert.Nil(t, err)
		assert.Equal(t, &ErrorResponse{Error: "Bad id"}, &errResp)
	})

	t.Run("service error", func(t *testing.T) {
		serviceMock := &PortServiceMock{
			GetPortFn: func(_ context.Context, in *pb.GetPortRequest, _ ...grpc.CallOption) (*pb.Port, error) {
				return nil, errors.New("some msg")
			},
		}
		service := New(serviceMock)

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/ports?id=55", nil)
		service.Router.ServeHTTP(rec, req)
		resp := rec.Result()
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})
}
