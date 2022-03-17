package realms

import (
	"GOOauth/Utils"
	"GOOauth/myDB"
	"context"
	"github.com/uptrace/bun"
	"log"
)

type realService interface {
	CreateOneInDb() (*realm, error)
}

func (r realm) CreateOneInDb() (*realm, error) {

	ctx, db, err := r.createTable()
	result, err := db.NewInsert().Model(&r).Exec(ctx)

	log.Println("result : ", result)

	if err != nil {
		return &realm{}, err
	}
	return &r, nil
}

func (r realm) createTable() (context.Context, *bun.DB, error) {
	ctx := context.Background()
	db := myDB.InitDb()
	_, err := db.NewCreateTable().Model((*realm)(nil)).IfNotExists().Exec(ctx)
	Utils.CheckAndWarnWfName(err, "createTable")

	return ctx, db, err
}
