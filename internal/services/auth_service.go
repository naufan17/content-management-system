package services

import (
	"github.com/naufan17/content-management-system/internal/dtos"
	"github.com/naufan17/content-management-system/internal/repositories"
	"github.com/naufan17/content-management-system/pkg/auth"

	"errors"
)

func LoginUser(user dtos.LoginDto) (dtos.AccessTokenDto, error) {
	userFromDB, err := repositories.GetUserByUsername(user.Username)

	if err != nil {
		return dtos.AccessTokenDto{}, errors.New("not found")
	}

	if !auth.ComparePassword(user.Password, userFromDB.Password) {
		return dtos.AccessTokenDto{}, errors.New("unauthorized")
	}

	accessAccessToken, accessExpiresIn, accessTokenType, err := auth.GenerateJWTAccess(userFromDB.ID)

	if err != nil {
		return dtos.AccessTokenDto{}, errors.New("internal server error")
	}

	return dtos.AccessTokenDto{
		AccessToken: accessAccessToken,
		ExpiresIn:   accessExpiresIn,
		TokenType:   accessTokenType,
	}, nil
}
