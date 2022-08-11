package db

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func New(dsn string) *bun.DB {
	postgresdb := sql.OpenDB((pgdriver.NewConnector(pgdriver.WithDSN(dsn))))
	return bun.NewDB(postgresdb, pgdialect.New())
}
