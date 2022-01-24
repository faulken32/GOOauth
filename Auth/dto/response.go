package dto

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// New
// create access token
func New(accessToken string, refreshToken string) AuthResponse {

	return AuthResponse{
		accessToken, refreshToken,
	}

}
