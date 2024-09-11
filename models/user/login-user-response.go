package user

import "time"

type LoginUserResponse struct {
	IsSuccess bool      `json:"isSuccess"`
	LoginAt   time.Time `json:"loginAt"`
}
