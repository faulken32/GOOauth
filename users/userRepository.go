package users

import (
	"GOOauth/Error"
	"GOOauth/Utils"
	"GOOauth/myDB"
	"context"
	"errors"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type QueryRes struct {
	Id   int
	Name string
	Uri  string
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

	salt, err := hashAndSalt(u.Password)
	if err != nil {
		return nil, Error.NewUserError(err)
	}
	u.Password = salt

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

func (u User) GetUserRealm() ([]QueryRes, error) {

	var query = "SELECT ep.id, ep.name  , ep.uri from \"user\" as u " +
		" inner join realms_users ru on u.id = ru.user_id " +
		" inner join end_points ep on ep.id = ru.realm_id" +
		" where u.login = ?  ;"
	db := myDB.InitDb()

	var items []QueryRes

	rows, err := db.Query(query, u.Login)

	if err != nil {
		return []QueryRes{}, err
	}

	defer func(db *bun.DB) {
		_ = db.Close()
	}(db)

	for rows.Next() {
		var r QueryRes
		err := rows.Scan(&r.Id, &r.Name, &r.Uri)
		if err != nil {
			return []QueryRes{}, err
		}
		items = append(items, r)
	}

	return items, nil
}

func hashAndSalt(pwd string) (string, error) {

	if pwd == "" {
		return "", errors.New("no password provided")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash), nil
}
