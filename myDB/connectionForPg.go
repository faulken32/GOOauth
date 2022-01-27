package myDB

import (
	"github.com/go-pg/pg"
)

func Connect() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "password",
	})

	return db
}

func Close(db *pg.DB) {

	defer db.Close()
}
