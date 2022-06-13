package users

import (
	"GOOauth/Error"
	"GOOauth/Utils"
	"GOOauth/myDB"
	"context"
	"github.com/uptrace/bun"
	"log"
)

type QueryRes struct {
	Id   string
	Name string
}

func (u User) TruncateTable() {
	log.Println("prepare test for user crud ----- TRUNCATE USER TABLE -------")
	//user := users.User{}
	//	realm := realms.Realm{}

	db := myDB.InitDb()

	_, err := db.NewTruncateTable().Model(&u).Exec(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

}

// CreateOne
// create one user in db
func (u User) CreateOne() (*User, *Error.UserError) {

	ctx, db, err := u.createTable()

	result, err := db.NewInsert().Model(&u).Exec(ctx)
	if err != nil {

		return nil, Error.NewUserError(err)
	}

	log.Println("created user table", result)
	Utils.CheckAndWarn(err)

	return &u, nil
}

func (u User) GetOneByLogin() *User {

	ctx := context.Background()
	db := myDB.InitDb()
	userDb := &User{}
	err := db.NewSelect().Model((*User)(nil)).Where("login = ?", u.Login).Scan(ctx, userDb)
	Utils.CheckAndWarn(err)
	return userDb
}

func (u User) createTable() (context.Context, *bun.DB, error) {
	ctx := context.Background()
	db := myDB.InitDb()
	_, err := db.NewCreateTable().Model((*User)(nil)).IfNotExists().Exec(ctx)
	Utils.CheckAndWarn(err)

	return ctx, db, err
}

func (u User) GetUserRealm() (QueryRes, error) {

	var query = "SELECT r.name from \"user\" as u " +
		" inner join realms_users ru on u.id = ru.user_id " +
		" inner join realms r on r.id = ru.realm_id" +
		" where u.login = ?  ;"
	db := myDB.InitDb()
	res := QueryRes{}

	var realmName string

	err := db.QueryRow(query, u.Login).Scan()

	if err != nil {
		return res, err
	}

	return res, nil
}