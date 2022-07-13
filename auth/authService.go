/*
 * Copyright (c) 2022.
 *
 * canicatti.nicolas@gmail.com
 *
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package auth

import (
	"GOOauth/Utils"
	"GOOauth/auth/dto"
	"GOOauth/users"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
)

type JWTMaker struct {
	secretKey string
}

type customClaims struct {
	Username   string           `json:"username"`
	Password   string           `json:"password"`
	Login      string           `json:"login"`
	AuthRealms []users.QueryRes `json:"authRealms"`
	jwt.StandardClaims
}

// DecodeAndValidateToken token validation
func DecodeAndValidateToken(token string) (bool, error, jwt.MapClaims) {
	secret := Utils.AppConfig.Secret

	maker := JWTMaker{secretKey: secret.SecretKey}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, nil
		}
		return []byte(maker.secretKey), nil
	}
	parse, err := jwt.Parse(token, keyFunc)
	claims := parse.Claims.(jwt.MapClaims)
	log.Println(claims)
	if err != nil {
		return false, err, nil
	}
	if parse.Valid {
		return true, nil, claims
	}
	return false, err, nil
}

// ExtractToken Extract token from request
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	return bearToken
}

// Authenticate  a user
func Authenticate(request *http.Request) (dto.Response, dto.ErrorResponse) {

	authRequest, err := decodeRequest(request)
	if err != nil {
		return dto.Response{}, dto.ErrorResponse{}
	}

	validatorError := authRequest.Validator(authRequest)

	if validatorError != nil {
		return dto.Response{}, dto.ErrorResponse{
			HttpStatus:   http.StatusBadRequest,
			ErrorMessage: validatorError.InternalErrorMessage,
		}
	}

	// check user in db
	fromRequest := users.NewFromRequest(authRequest)
	auth := fromRequest.MapToUser().Auth(fromRequest.EndPoint)

	if err != nil {
		return dto.Response{}, dto.ErrorResponse{
			HttpStatus:   500,
			ErrorMessage: err.Error(),
		}
	}
	if auth {

		var response dto.Response
		if auth {
			return encodeAuthResponse(authRequest, response)
		} else {
			return dto.Response{}, dto.ErrorResponse{
				HttpStatus:   403,
				ErrorMessage: "You don't have right on realm",
			}
		}

	}

	if err != nil {
		return dto.Response{}, dto.ErrorResponse{}
	}

	return dto.Response{}, dto.ErrorResponse{
		HttpStatus:   403,
		ErrorMessage: "You don't have right on realm",
	}
}

func decodeRequest(request *http.Request) (dto.AuthRequest, error) {
	var authRequest dto.AuthRequest

	j := json.NewDecoder(request.Body)
	err := j.Decode(&authRequest)
	return authRequest, err
}
