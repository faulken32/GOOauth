package users

import (
	"GOOauth/Utils"
	"GOOauth/myDB"
	"context"
	"encoding/json"
	"github.com/uptrace/bun"
	"log"
)

type UserService interface {
	CreateOne() (*UserDb, *UserError)
}

// UserDb
// Base model entity
type UserDb struct {
	bun.BaseModel `bun:",table:user"`
	ID            int64  `bun:"id,pk,autoincrement"`
	Login         string `bun:"login,unique"`
	Name          string `bun:"name,unique"`
	Email         string `bun:"email,unique"`
}

func NewUserDbNoId(login string, name string, email string) *UserDb {
	return &UserDb{Login: login, Name: name, Email: email}
}
func (u UserDb) String() {
	log.Printf("%d %s %s %s", u.ID, u.Login, u.Name, u.Email)
}

// export struct to json string
func (u UserDb) ToJson() {
	marshal, err := json.Marshal(u)
	Utils.CheckAndWarn(err)
	log.Println(string(marshal))

}

// CreateOne
// create one user in db
func (u UserDb) CreateOne() (*UserDb, *UserError) {

	ctx, db, err := u.createTable()

	result, err := db.NewInsert().Model(&u).Exec(ctx)
	if err != nil {
		log.Println(err)
		return nil, NewUserError("error from db", err)
	}

	log.Println("created user table", result)
	Utils.CheckAndWarn(err)

	return &u, nil
}

func GetOneByLogin(login string) *UserDb {

	ctx := context.Background()
	db := myDB.InitDb()
	userDb := &UserDb{}
	err := db.NewSelect().Model((*UserDb)(nil)).Where("login = ?", login).Scan(ctx, userDb)
	Utils.CheckAndWarn(err)
	return userDb
}

func (u UserDb) createTable() (context.Context, *bun.DB, error) {
	ctx := context.Background()
	db := myDB.InitDb()
	_, err := db.NewCreateTable().Model((*UserDb)(nil)).IfNotExists().Exec(ctx)
	Utils.CheckAndWarn(err)

	return ctx, db, err
}
