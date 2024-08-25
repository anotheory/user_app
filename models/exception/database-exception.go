package exception

type DatabaseException struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

func NewDatabaseException(message string) *DatabaseException {
	return &DatabaseException{
		ErrorCode:    "DATABASE_EXCEPTION",
		ErrorMessage: message,
	}
}

func (err *DatabaseException) Error() string {
	return err.ErrorMessage
}
