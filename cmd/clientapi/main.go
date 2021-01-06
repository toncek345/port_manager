package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/toncek345/port-manager/internal/clientapi"
)

func main() {
	// svc addr

	// graceful shutdown

	port, err := strconv.ParseInt(os.Getenv("PORT"), 10, 32)
	if err != nil {
		log.Fatalf("port cannot be parsed: %s", err)
		return
	}

	api := clientapi.New(int(port))

	go func() {
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
