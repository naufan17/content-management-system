package dto

import "github.com/naufan17/content-management-system/internal/model"

type UserDto struct {
	Name string `json:"name"`
}

func UserModelToDto(user model.User) UserDto {
	return UserDto{
		Name: user.Name,
	}
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,max=50"`
	Password string `json:"password" validate:"required,min=10"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"`
}
