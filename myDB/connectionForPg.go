package myDB

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func InitDb() *bun.DB {

	dsn := "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"
	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqlDb, pgdialect.New())

	return db
}
