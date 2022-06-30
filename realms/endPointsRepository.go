package realms

import (
	"GOOauth/Utils"
	"GOOauth/myDB"
	"context"
	"errors"
	"github.com/uptrace/bun"
	"log"
)

type EndPointsRepository interface {
	FindById() (*Endpoint, error)
	FindByUrl() (*Endpoint, error)
}

func (e Endpoint) FindById() (*Endpoint, error) {
	model := &Endpoint{}
	if e.ID != 0 {
		db := myDB.InitDb()
		err := db.NewSelect().Model(model).Where("id = ?", e.ID).Scan(context.Background(), model)
		if err != nil {
			return &Endpoint{}, err
		}
		return model, nil
	}

	return model, errors.New("missing id")
}

func (e Endpoint) FindByUrl() (*Endpoint, error) {
	model := &Endpoint{}
	if e.Url != "" {
		db := myDB.InitDb()
		err := db.NewSelect().Model(model).Where("url = ?", e.ID).Scan(context.Background(), model)
		if err != nil {
			return &Endpoint{}, err
		}
		return model, nil
	}

	return model, errors.New("missing url")
}

func (e Endpoint) FindByUri() (*Endpoint, error) {
	model := &Endpoint{}
	if e.Uri != "" {
		db := myDB.InitDb()
		err := db.NewSelect().Model(model).Where("uri = ?", e.Uri).Scan(context.Background(), model)
		if err != nil {
			return &Endpoint{}, err
		}
		return model, nil
	}

	return model, errors.New("missing uri")
}

func (e Endpoint) TruncateTable() {
	log.Println("prepare test for endpoints crud ----- TRUNCATE endpoints TABLE -------")

	db := myDB.InitDb()

	_, err := db.NewTruncateTable().Model(&e).Exec(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

}

func (e Endpoint) Save() (*Endpoint, error) {

	ctx, db, err := e.createTable()
	result, err := db.NewInsert().Model(&e).Exec(ctx)

	log.Println("result : ", result)

	if err != nil {
		return &Endpoint{}, err
	}
	return &e, nil
}

func (e Endpoint) createTable() (context.Context, *bun.DB, error) {
	ctx := context.Background()
	db := myDB.InitDb()
	_, err := db.NewCreateTable().Model((*Endpoint)(nil)).IfNotExists().Exec(ctx)
	Utils.CheckAndWarnWfName(err, "createTable")

	return ctx, db, err
}
