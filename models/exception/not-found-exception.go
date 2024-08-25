package exception

type NotFoundException struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

func NewNotFoundException(message string) *NotFoundException {
	return &NotFoundException{
		ErrorCode:    "NOT_FOUND",
		ErrorMessage: message,
	}
}

func (err *NotFoundException) Error() string {
	return err.ErrorMessage
}
