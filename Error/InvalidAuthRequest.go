package Error

type InvalidAuthRequest struct {
	InternalErrorCode    int
	InternalErrorMessage string
}

const (
	MissingLogin    = "missing login"
	MissingPassword = "missing password"
	MissingRealm    = "missing realm"
)

func NewInvalidAuthRequest(internalErrorCode int, internalErrorMessage string) *InvalidAuthRequest {
	return &InvalidAuthRequest{InternalErrorCode: internalErrorCode, InternalErrorMessage: internalErrorMessage}
}
