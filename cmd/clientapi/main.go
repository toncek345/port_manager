package main

import "github.com/toncek345/port-manager/internal/clientapi"

func main() {
	// port
	// svc addr

	// graceful shutdown

	api := clientapi.New(5000)
	api.PortsHandler(nil, nil)
}
