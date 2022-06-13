package users

import (
	"GOOauth/Error"
	"encoding/json"
	"github.com/uptrace/bun"
	"log"
)

type (
	// User
	// Base model entity
	User struct {
		bun.BaseModel `bun:",table:user"`
		ID            int64  `bun:"id,pk,autoincrement"`
		Login         string `bun:"login,unique"`
		Name          string `bun:"name,unique"`
		Email         string `bun:"email,unique"`
		Password      string `bun:"password"`
	}

	UserObj interface {
		ToJson()
		ToString()
	}

	UserRepository interface {
		CreateOne() (*User, *Error.UserError)
		GetOneByLogin() *User
		GetUserRealm() string
		TruncateTable()
	}
)

func NewUser(login string, name string, email string, password string) *User {
	return &User{Login: login, Name: name, Email: email, Password: password}
}

// export struct to json string
func (u User) ToJson() {
	marshal, _ := json.Marshal(u)
	//Utils.CheckAndWarn(err)
	log.Println(string(marshal))

}

func (u User) ToString() {
	log.Printf("%d %s %s %s", u.ID, u.Login, u.Name, u.Email)
}
