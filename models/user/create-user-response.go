package user

import (
	"time"
	"user_app/models/base"
)

type CreateUserResponse struct {
	Username  string        `json:"username"`
	Email     string        `json:"email"`
	IsSuccess base.NullBool `json:"isSuccess"`
	CreatedAt time.Time     `json:"createdAt"`
}
