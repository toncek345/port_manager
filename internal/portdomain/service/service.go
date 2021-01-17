package service

import "github.com/jmoiron/sqlx"

type Services struct {
	Port PortService
}

func New(db *sqlx.DB) *Services {
	return &Services{
		Port: &PortSQL{
			db: db,
		},
	}
}
