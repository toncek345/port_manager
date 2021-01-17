package clientapi

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	pb "github.com/toncek345/port_manager/internal/portdomain/proto"
)

func (c *clientAPI) upsertPorts(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)

	_, err := dec.Token()
	if err != nil {
		c.JSON(w, http.StatusBadRequest, ErrorResponse{Error: "Bad json"})
		return
	}

	upsert, err := c.serviceClient.Upsert(r.Context())
	if err != nil {
		c.JSON(w, http.StatusInternalServerError, err)
		return
	}
	defer upsert.CloseAndRecv()

	for dec.More() {
		portID, err := dec.Token()
		if err != nil {
			panic(err)
		}

		port := &Port{}
		if err = dec.Decode(port); err != nil {
			c.JSON(w, http.StatusInternalServerError, err)
			return
		}

		portIDStr, ok := portID.(string)
		if !ok {
			c.JSON(w, http.StatusBadRequest, ErrorResponse{Error: "Bad json"})
			return
		}

		port.IDStr = portIDStr

		if err := upsert.Send(PortToPortProto(port)); err != nil {
			c.JSON(w, http.StatusInternalServerError, err)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *clientAPI) getPort(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		c.JSON(w, http.StatusBadRequest, ErrorResponse{Error: "Bad id"})
		return
	}

	port, err := c.serviceClient.GetPort(r.Context(), &pb.GetPortRequest{PortId: id})
	if err != nil {
		c.JSON(w, http.StatusInternalServerError, err)
		return
	}

	data, err := json.Marshal(PortProtoToPort(port))
	if err != nil {
		c.JSON(w, http.StatusInternalServerError, err)
		return
	}

	w.Write(data)
}

func (c *clientAPI) PortsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.getPort(w, r)
	case http.MethodPost:
		c.upsertPorts(w, r)
	default:
		log.Println("wrong method")
		c.JSON(w, http.StatusBadRequest, &ErrorResponse{Error: "Wrong request method"})
		return
	}
}
