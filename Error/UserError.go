package Error

import (
	"GOOauth/Utils"
	"encoding/json"
	"log"
)

type UserError struct {
	Err error
}

func NewUserError(err error) *UserError {
	return &UserError{Err: err}

}

func (u UserError) Error() string {

	return u.Err.Error()
}

func (u UserError) ToJson() {
	marshal, err := json.Marshal(u.Err.Error())
	Utils.CheckAndWarn(err)
	log.Println(string(marshal))

}
