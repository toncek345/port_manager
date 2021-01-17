package port

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Service interface {
	UpsertPort(ctx context.Context, port *Port) error
	GetPort(ctx context.Context, id int64) (*Port, error)
}

var _ Service = (*SQL)(nil)

type SQL struct {
	db *sqlx.DB
}

func NewSQL(db *sqlx.DB) *SQL {
	return &SQL{
		db: db,
	}
}

func (p *SQL) GetPort(ctx context.Context, id int64) (*Port, error) {
	var port Port

	if err := p.db.GetContext(
		ctx,
		&port,
		"SELECT * FROM ports WHERE id = $1",
		id); err != nil {
		return nil, fmt.Errorf("getting port: %w", err)
	}

	return &port, nil
}

func (p *SQL) UpsertPort(ctx context.Context, port *Port) error {
	var pnew Port

	if err := p.db.GetContext(
		ctx,
		&pnew,
		"SELECT * FROM ports WHERE id_str = $1",
		port.IDStr,
	); err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("getting port: %w", err)
	}

	if pnew.ID == 0 {
		return p.insertPort(ctx, port)
	}

	if _, err := p.db.Exec(
		`UPDATE ports SET id_str=$1, name=$2, city=$3, country=$4, coord_long=$5, coord_lat=$6,
		province=$7, timezone=$8, code=$9, regions=$10, unlocs=$11, alias=$12 WHERE id = $13`,
		port.IDStr, port.Name, port.City, port.Country, port.CoordinatesLon, port.CoordinatesLat,
		port.Provice, port.Timezone, port.Code, port.Regions, port.Unlocs, port.Alias,
		port.ID,
	); err != nil {
		return fmt.Errorf("inserting port: %w", err)
	}

	return nil
}

func (p *SQL) insertPort(ctx context.Context, port *Port) error {
	if _, err := p.db.Exec(
		`INSERT INTO ports (id_str, name, city, country, coord_long, coord_lat,
		province, timezone, code, regions, unlocs, alias) VALUES
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
		port.IDStr, port.Name, port.City, port.Country, port.CoordinatesLon, port.CoordinatesLat,
		port.Provice, port.Timezone, port.Code, port.Regions, port.Unlocs, port.Alias,
	); err != nil {
		return fmt.Errorf("inserting port: %w", err)
	}

	return nil
}
