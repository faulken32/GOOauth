package dto

type Response struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// NewResponse
// create access token
func NewResponse(accessToken string, refreshToken string) Response {

	return Response{
		accessToken, refreshToken,
	}

}
