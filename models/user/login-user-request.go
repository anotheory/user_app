package user

type LoginUserRequest struct {
	Username string `json:"username" validate:"omitempty,min=8,max=50"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"sha256"`
}
