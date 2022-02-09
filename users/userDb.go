package users

import (
	"GOOauth/Utils"
	"GOOauth/myDB"
	"context"
	"github.com/uptrace/bun"
	"log"
)

type UserDb struct {
	bun.BaseModel `bun:",table:user"`
	ID            int64  `bun:"id,pk,autoincrement"`
	Login         string `bun:"login,unique"`
	Name          string `bun:"name,unique"`
	Email         string `bun:"email,unique"`
}

func NewUserDb(id int64, login string, name string, email string) *UserDb {
	return &UserDb{ID: id, Login: login, Name: name, Email: email}
}

func NewUserDbNoId(login string, name string, email string) *UserDb {
	return &UserDb{Login: login, Name: name, Email: email}
}

func (u UserDb) CreateOne() *UserDb {

	ctx := context.Background()
	db := myDB.InitDb()
	exec, err := db.NewCreateTable().Model((*UserDb)(nil)).IfNotExists().Exec(ctx)
	Utils.CheckAndWarn(err)
	log.Println("created table { }", exec)

	result, err := db.NewInsert().Model(&u).Exec(ctx)
	Utils.CheckAndWarn(err)
	log.Println("created user { }", result)
	db.Close()
	//err := connection.CreateTable(&u, &orm.CreateTableOptions{
	//	IfNotExists: true,
	//})
	//Utils.CheckAndWarn(err)
	//err = connection.Insert(&u)
	//
	//Utils.CheckAndWarn(err)
	//
	//myDB.Close(connection)
	return &u
}
