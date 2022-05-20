package dto

import (
	"GOOauth/Error"
)

type RequestValidator interface {
	Validator(request AuthRequest) *Error.InvalidAuthRequest
}

type AuthRequest struct {
	Login    string
	Password string
	Realm    string
}

func (r AuthRequest) Validator(request AuthRequest) *Error.InvalidAuthRequest {
	if request.Login == "" {

		return Error.NewInvalidAuthRequest(0, Error.MissingLogin)
	}

	if request.Password == "" {

		return Error.NewInvalidAuthRequest(1, Error.MissingPassword)
	}

	if request.Realm == "" {

		return Error.NewInvalidAuthRequest(3, Error.MissingRealm)
	}
	return nil
}

func NewRequest(login string, password string, realm string) *AuthRequest {
	return &AuthRequest{Login: login, Password: password, Realm: realm}
}
