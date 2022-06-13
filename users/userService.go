package users

import (
	"GOOauth/Error"
	"errors"
)

type UserService interface {
	ValidateIdentity(password string) (bool, *Error.UserError)
}

func (u User) ValidateIdentity(password string) (bool, *Error.UserError) {

	if password == "" {
		err := errors.New("no password provided")
		return false, Error.NewUserError(err)
	}
	if u.Password == password {
		return true, nil
	}
	return false, nil
}
