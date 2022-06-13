package myDB

import (
	"GOOauth/users"
	"context"
	"log"
	"testing"
)

func TestConnect(t *testing.T) {

	initDb := InitDb()
	ctx := context.Background()

	// The request has a timeout, so create a context that is
	// canceled automatically when the timeout expires.

	res, _ := initDb.ExecContext(ctx, "SELECT 1")
	//Utils.CheckAndWarn(err)
	log.Println(res)

}

func TestTruncateTable(t *testing.T) {

	truncateTable(users.User{})
}
