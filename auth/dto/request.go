package dto

import "GOOauth/Error"

type RequestValidator interface {
	Validator(request Request) *Error.InvalidAuthRequest
}

type Request struct {
	Login    string
	Password string
	Realm    string
}

func (r Request) Validator(request Request) *Error.InvalidAuthRequest {
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

func NewRequest(login string, password string, realm string) *Request {
	return &Request{Login: login, Password: password, Realm: realm}
}
