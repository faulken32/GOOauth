package proxy

import (
	"GOOauth/auth"
	"GOOauth/realms"
	"github.com/golang-jwt/jwt"
	"io"
	"log"
	"net/http"
)

func Main(w http.ResponseWriter, r *http.Request) {
	uri := r.URL.RequestURI()

	token := auth.ExtractToken(r)
	if token == "" {
		w.WriteHeader(403)
		w.Write([]byte("no auth token"))
		return
	}

	validateToken, _, claims := auth.DecodeAndValidateToken(token)
	if !validateToken {
		w.WriteHeader(403)
		w.Write([]byte("token invalid"))
		return
	}
	authorization := gotAuthorization(claims, uri)
	if !authorization {
		w.WriteHeader(403)
		w.Write([]byte("invalid authorization"))
		return
	}

	endpoint := &realms.Endpoint{Uri: uri}
	endpoint, err := endpoint.FindByUri()

	if err != nil {
		log.Println(err.Error())
		return
	}
	request, err := http.NewRequest(endpoint.Method, endpoint.Url, r.Body)
	client := &http.Client{}
	do, err := client.Do(request)
	response := readResponseFromHost(err, do)
	_, _ = w.Write(response)
	log.Println("end")
}

func readResponseFromHost(errFromTargetedHost error, resp *http.Response) []byte {

	if errFromTargetedHost != nil {
		log.Println(errFromTargetedHost.Error())
	}
	bodyBytes, errFromTargetedHost := io.ReadAll(resp.Body)
	return bodyBytes
}

func gotAuthorization(claims jwt.MapClaims, requestedUrl string) bool {

	for k, v := range claims {
		if k == "authRealms" {
			res, ok := v.(interface{})
			if ok {
				m := res.([]interface{})
				for _, i2 := range m {
					queryRes := i2.(map[string]interface{})
					for typeName, value := range queryRes {
						if typeName == "Uri" && value == requestedUrl {
							return true
						}
					}
				}
			}
			return false
		}
	}
	return false
}
