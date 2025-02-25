package model

type UserResponse struct {
	ID        string `json:"id,omitempty" example:"user1"`
	Name      string `json:"name,omitempty" example:"User 1"`
	Token     string `json:"token,omitempty" example:"01953b4e-78ac-73dc-8480-da4ec118ef28"`
	CreatedAt int64  `json:"created_at,omitempty" example:"1630000000"`
	UpdatedAt int64  `json:"updated_at,omitempty" example:"1630000000"`
}

type VerifyUserRequest struct {
	Token string `validate:"required,max=100" example:"01953b4e-78ac-73dc-8480-da4ec118ef28"`
}

type RegisterUserRequest struct {
	ID       string `json:"id" validate:"required,max=100" example:"user1"`
	Password string `json:"password" validate:"required,max=100" example:"password"`
	Name     string `json:"name" validate:"required,max=100" example:"User 1"`
}

type UpdateUserRequest struct {
	ID       string `json:"-" validate:"required,max=100" example:"user1"`
	Password string `json:"password,omitempty" validate:"max=100" example:"password"`
	Name     string `json:"name,omitempty" validate:"max=100" example:"User 1"`
}

type LoginUserRequest struct {
	ID       string `json:"id" validate:"required,max=100" example:"user1"`
	Password string `json:"password" validate:"required,max=100" example:"password"`
}

type LogoutUserRequest struct {
	ID string `json:"id" validate:"required,max=100" example:"user1"`
}

type GetUserRequest struct {
	ID string `json:"id" validate:"required,max=100" example:"user1"`
}
