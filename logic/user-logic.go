package logic

import (
	"database/sql"
	"regexp"
	"user_app/models/base"
	"user_app/models/entity"
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
		return exception.NewValidationException("Invalid request body.")
	} else if matched, _ = regexp.MatchString(validUsernamePattern, request.Username); !matched {
		return exception.NewValidationException("Wrong format on username.")
	} else if matched, _ = regexp.MatchString(invalidPasswordPattern, request.Password); matched {
		return exception.NewValidationException("Wrong format on password.")
	}
	return nil
}

func MapCreateUserResponse(entity entity.UserEntity) *user.CreateUserResponse {
	response := user.CreateUserResponse{
		Username: entity.Username,
		Email:    entity.Email,
		IsSuccess: base.NullBool{
			NullBool: sql.NullBool{
				Bool:  true,
				Valid: true,
			},
		},
		CreatedAt: entity.CreatedAt,
	}
	return &response
}

func ValidateLoginUserRequest(request user.LoginUserRequest) error {
	var err error
	validateFunc := validator.New(validator.WithRequiredStructEnabled())
	if err = validateFunc.Struct(request); err != nil {
		return exception.NewValidationException("Invalid request body.")
	} else if request.Username == "" && request.Email == "" {
		return exception.NewValidationException("Invalid request body.")
	} else if (request.Username != "") == (request.Email != "") {
		return exception.NewValidationException("Choose one between username and email.")
	}
	return nil
}

func MapLoginUserResponse(entity entity.UserLoginHistoryEntity) *user.LoginUserResponse {
	response := user.LoginUserResponse{
		IsSuccess: true,
		LoginAt:   entity.LoginAt,
	}
	return &response
}
