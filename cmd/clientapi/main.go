package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/toncek345/port_manager/internal/clientapi"
	pb "github.com/toncek345/port_manager/internal/portdomain/proto"
	"google.golang.org/grpc"
)

func main() {
	port, err := strconv.ParseInt(os.Getenv("PORT"), 10, 32)
	if err != nil {
		log.Fatalf("port cannot be parsed: %s\n", err.Error())
	}

	svcAddr := os.Getenv("SVC")
	if svcAddr == "" {
		log.Fatalln("wrong svcaddr")
	}

	svcConn, err := grpc.Dial(svcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("couldn't dial grpc service")
	}
	defer svcConn.Close()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: clientapi.New(pb.NewPortDomainClient(svcConn)).Router,
	}

	srvErrChan := make(chan error, 1)
	go func() {
		log.Printf("server started on port: %d\n", port)
		srvErrChan <- server.ListenAndServe()
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT)

	select {
	case <-signalChan:
		log.Println("shutting down...")
		server.Shutdown(context.Background())
	case err := <-srvErrChan:
		log.Fatalf("server error: %s\n", err.Error())
	}
}
