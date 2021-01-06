package clientapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type clientAPI struct {
	router *http.ServeMux
	port   int
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func New(port int) *clientAPI {
	api := &clientAPI{
		router: http.NewServeMux(),
		port:   port,
	}

	api.router.HandleFunc("/ports", api.PortsHandler)

	return api
}

// Start runs http server.
func (c *clientAPI) Start() {

}

// Stop stops http server.
func (c *clientAPI) Stop() {

}

func (c *clientAPI) JSON(w http.ResponseWriter, status int, obj interface{}) error {
	data, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("unable to marshall json: %w", err)
	}

	if status != http.StatusOK {
		w.WriteHeader(status)
	}
	w.Write(data)
	return nil
}

func (a *clientAPI) PortsHandler(w http.ResponseWriter, r *http.Request) {

}
