package db

import (
	"GOOauth/Utils"
	"GOOauth/myDB"
	"context"
	"database/sql"
	"github.com/uptrace/bun"
)

func CreateAndInsert(model *interface{}) (sql.Result, error) {
	ctx, db, _ := CreateTableIfNotExist(model)
	return insertInDb(db, ctx, model)
}

func CreateTableIfNotExist(model *interface{}) (context.Context, *bun.DB, error) {
	ctx := context.Background()
	db := myDB.InitDb()
	_, err := db.NewCreateTable().Model((*model)(nil)).IfNotExists().Exec(ctx)
	Utils.CheckAndWarnWfName(err, "createTable")
	myDB.Close(db)
	return ctx, db, err
}

func insertInDb(db *bun.DB, ctx context.Context, r interface{}) (sql.Result, error) {

	result, err := db.NewInsert().Model(&r).Exec(ctx)
	myDB.Close(db)
	return result, err
}
