package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/toncek345/port_manager/internal/portdomainsvc"
)

func main() {
	port, err := strconv.ParseInt(os.Getenv("PORT"), 10, 32)
	if err != nil {
		log.Fatalf("port cannot be parsed: %s", err)
		return
	}

	dbConnStr := os.Getenv("DB")
	if dbConnStr == "" {
		log.Fatal("invalid db conn str")
	}

	db, err := sqlx.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatal("opening db conn: %w", err)
	}

	service := portdomainsvc.New(int(port), db)

	go func() {
		// TODO: this is not ideal, server can die and the proccess will still be running
		if err := service.Start(); err != nil {
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

	service.Stop()
}
