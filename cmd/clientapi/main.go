package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/toncek345/port_manager/internal/clientapi"
)

func main() {
	port, err := strconv.ParseInt(os.Getenv("PORT"), 10, 32)
	if err != nil {
		log.Fatalf("port cannot be parsed: %s", err)
	}

	svcAddr := os.Getenv("SVC")
	if svcAddr == "" {
		log.Fatal("wrong svcaddr")
	}

	api, err := clientapi.New(int(port), svcAddr)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		// TODO: this is not ideal, server can die and the proccess will still be running
		if err := api.Start(); err != nil {
			log.Fatal(err)
			return
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT)

	<-signalChan
	log.Print("shutting down...\n")

	api.Stop()
}
