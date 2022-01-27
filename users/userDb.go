package users

import (
	"GOOauth/Utils"
	"GOOauth/myDB"
	"github.com/go-pg/pg/orm"
)

type UserDb struct {
	Id    int64
	Login string
	Name  string
	Email string
}

func NewUserDb(id int64, login string, name string, email string) *UserDb {
	return &UserDb{Id: id, Login: login, Name: name, Email: email}
}

func NewUserDbNoId(login string, name string, email string) *UserDb {
	return &UserDb{Login: login, Name: name, Email: email}
}

func (u UserDb) CreateOne() *UserDb {
	connection := myDB.Connect()

	err := connection.CreateTable(&u, &orm.CreateTableOptions{
		IfNotExists: true,
	})
	Utils.CheckAndWarn(err)
	err = connection.Insert(&u)

	Utils.CheckAndWarn(err)

	myDB.Close(connection)
	return &u
}
