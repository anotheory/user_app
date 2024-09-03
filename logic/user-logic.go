package logic

import (
	"regexp"
	"user_app/models/exception"
	"user_app/models/user"

	"github.com/go-playground/validator/v10"
)

var validUsernamePattern string = "^[a-zA-Z]([\\w]{6,48})[a-zA-Z0-9]$"
var invalidPasswordPattern string = "\\s+"

func ValidateCreateUserRequest(request user.CreateUserRequest) error {
	var err error
	var matched bool = false
	validateFunc := validator.New(validator.WithRequiredStructEnabled())
	if err = validateFunc.Struct(request); err != nil {
		err = exception.NewValidationException("Malformed request body.")
		return err
	} else if matched, _ = regexp.MatchString(validUsernamePattern, request.Username); !matched {
		err = exception.NewValidationException("Wrong format on username.")
		return err
	} else if matched, _ = regexp.MatchString(invalidPasswordPattern, request.Password); matched {
		err = exception.NewValidationException("Wrong format on password.")
		return err
	}
	return nil
}
