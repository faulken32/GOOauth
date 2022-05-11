package myDB

import (
	"GOOauth/Utils"
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func InitDb() *bun.DB {

	dsn := "postgres://postgres:toto@localhost:5432/postgres?sslmode=disable"
	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqlDb, pgdialect.New())

	return db
}

func Close(db *bun.DB) {

	defer func(db *bun.DB) {
		err := db.Close()
		Utils.CheckAndWarn(err)
	}(db)
}
