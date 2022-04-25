package realms

import (
	"GOOauth/Utils"
	"GOOauth/myDB"
	"context"
	"github.com/uptrace/bun"
	"log"
)

type RealService interface {
	CreateOneInDb() (*Realm, error)
}

func (r Realm) CreateOneInDb() (*Realm, error) {

	ctx, db, err := r.createTable()
	result, err := db.NewInsert().Model(&r).Exec(ctx)

	log.Println("result : ", result)

	if err != nil {
		return &Realm{}, err
	}
	return &r, nil
}

func (r Realm) createTable() (context.Context, *bun.DB, error) {
	ctx := context.Background()
	db := myDB.InitDb()
	_, err := db.NewCreateTable().Model((*Realm)(nil)).IfNotExists().Exec(ctx)
	Utils.CheckAndWarnWfName(err, "createTable")

	return ctx, db, err
}
