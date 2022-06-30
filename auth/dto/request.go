package dto

import (
	"GOOauth/Error"
)

type RequestValidator interface {
	Validator(request AuthRequest) *Error.InvalidAuthRequest
}

type AuthRequest struct {
	Login             string
	Name              string
	Password          string
	RequestedEndPoint string
}

func NewAuthRequest(login string, name string, password string) *AuthRequest {
	return &AuthRequest{Login: login, Name: name, Password: password}
}

func (r AuthRequest) Validator(request AuthRequest) *Error.InvalidAuthRequest {
	if request.Login == "" {

		return Error.NewInvalidAuthRequest(0, Error.MissingLogin)
	}

	if request.Password == "" {

		return Error.NewInvalidAuthRequest(1, Error.MissingPassword)
	}

	return nil
}
