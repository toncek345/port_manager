package clientapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	pb "github.com/toncek345/port_manager/internal/portdomainsvc/grpc"
	"google.golang.org/grpc"
)

type clientAPI struct {
	router *http.ServeMux
	port   int

	serviceClient pb.PortDomainClient

	server *http.Server
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func New(port int, grpcServiceAddr string) (*clientAPI, error) {
	conn, err := grpc.Dial(grpcServiceAddr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("cannot connect to grpc server: %w", err)
	}

	api := &clientAPI{
		router:        http.NewServeMux(),
		port:          port,
		serviceClient: pb.NewPortDomainClient(conn),
	}

	api.router.HandleFunc("/ports", api.PortsHandler)

	return api, nil
}

// Start runs http server.
func (c *clientAPI) Start() error {
	c.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", c.port),
		Handler: c.router,
	}

	if err := c.server.ListenAndServe(); err != nil {
		return fmt.Errorf("listening and serving: %w", err)
	}

	return nil
}

// Stop stops http server.
func (c *clientAPI) Stop() error {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	if err := c.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutting down server: %w", err)
	}

	return nil
}

func (c *clientAPI) JSON(w http.ResponseWriter, status int, obj interface{}) error {
	data, err := json.Marshal(obj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return fmt.Errorf("unable to marshall json: %w", err)
	}

	if status != http.StatusOK {
		w.WriteHeader(status)
	}
	w.Write(data)
	return nil
}
