package service

import (
	"errors"

	"github.com/naufan17/content-management-system/internal/dto"
	"github.com/naufan17/content-management-system/internal/repository"
	"github.com/naufan17/content-management-system/pkg/utils"
)

func LoginUser(user dto.LoginRequest) (dto.LoginResponse, error) {
	userFromDB, err := repository.FindByUsernameUser(user.Username)

	if err != nil {
		return dto.LoginResponse{}, errors.New("not found")
	}

	if !utils.ComparePassword(user.Password, userFromDB.Password) {
		return dto.LoginResponse{}, errors.New("unauthorized")
	}

	accessAccessToken, accessExpiresIn, accessTokenType, err := utils.GenerateJWT(userFromDB.ID)

	if err != nil {
		return dto.LoginResponse{}, errors.New("internal server error")
	}

	return dto.LoginResponse{
		AccessToken: accessAccessToken,
		ExpiresIn:   accessExpiresIn,
		TokenType:   accessTokenType,
	}, nil
}
