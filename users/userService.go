package users

import (
	"GOOauth/Error"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserService interface {
	ValidateIdentity(password string) (bool, *Error.UserError)
}

func (u User) ValidateIdentity(passwordToValidate string) (bool, *Error.UserError) {

	if passwordToValidate == "" {
		err := errors.New("no passwordToValidate provided")
		return false, Error.NewUserError(err)
	}

	byteHash := []byte(u.Password)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(passwordToValidate))
	if err != nil {
		log.Println(err)
		return false, Error.NewUserError(err)
	}

	return true, nil
}
