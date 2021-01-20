package clientapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	pb "github.com/toncek345/port_manager/internal/portdomain/proto"
)

type clientAPI struct {
	Router        *http.ServeMux
	serviceClient pb.PortDomainClient
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func New(portDomainSvcClient pb.PortDomainClient) *clientAPI {
	if portDomainSvcClient == nil {
		log.Fatalln("port domain client is required")
	}

	api := &clientAPI{
		Router:        http.NewServeMux(),
		serviceClient: portDomainSvcClient,
	}

	api.Router.HandleFunc("/ports", api.PortsHandler)

	return api
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
