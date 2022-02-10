package users

import (
	"GOOauth/Utils"
	"encoding/json"
	"log"
)

type UserError struct {
	Message string
	Error   error
}

func NewUserError(message string, error error) *UserError {
	return &UserError{Message: message, Error: error}
}

func (u UserError) ToJson() {
	marshal, err := json.Marshal(u.Error.Error())
	Utils.CheckAndWarn(err)
	log.Println(string(marshal))

}
