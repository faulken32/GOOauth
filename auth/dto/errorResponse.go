package dto

type ErrorResponse struct {
	HttpStatus   int
	ErrorMessage string `json:"ErrorMessage"`
}

func New(httpStatus int, errorMessage string) ErrorResponse {

	return ErrorResponse{httpStatus, errorMessage}

}
