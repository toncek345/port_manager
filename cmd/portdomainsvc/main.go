package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/toncek345/port_manager/internal/portdomain/server"
	"github.com/toncek345/port_manager/internal/portdomain/service"
)

func main() {
	port, err := strconv.ParseInt(os.Getenv("PORT"), 10, 32)
	if err != nil {
		log.Fatalf("port cannot be parsed: %s", err.Error())
		return
	}

	dbConnStr := os.Getenv("DB")
	if dbConnStr == "" {
		log.Fatalln("invalid db conn str")
	}

	db, err := sqlx.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalln("opening db conn: %w", err)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("listening failed: %s\n", err.Error())
	}

	server := server.New(service.New(db))

	srvErrChan := make(chan error, 1)
	go func() {
		log.Printf("starting server on port %d\n", port)
		srvErrChan <- server.Serve(listener)
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT)

	select {
	case <-signalChan:
		log.Println("shutting down...")
		server.Close()
	case err := <-srvErrChan:
		log.Fatalf("server error: %s\n", err.Error())
	}
}
