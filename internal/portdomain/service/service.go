package service

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/toncek345/port_manager/internal/portdomain/service/port"
)

type Services struct {
	Port port.Service
}

func New(db *sqlx.DB) *Services {
	if db == nil {
		log.Fatalln("db is required")
	}

	return &Services{
		Port: port.NewSQL(db),
	}
}
