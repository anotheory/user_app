package exception

type ValidationException struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

func NewValidationException(message string) *ValidationException {
	return &ValidationException{
		ErrorCode:    "BAD_REQUEST",
		ErrorMessage: message,
	}
}

func (err *ValidationException) Error() string {
	return err.ErrorMessage
}
