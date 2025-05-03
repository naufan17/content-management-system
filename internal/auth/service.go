package auth

import (
	"errors"

	"github.com/naufan17/content-management-system/pkg/utils"
)

type AuthService interface {
	LoginUser(user LoginRequest) (LoginResponse, error)
}

type authService struct {
	userRepository UserRepository
}

func NewAuthService(userRepository UserRepository) AuthService {
	return &authService{
		userRepository: userRepository,
	}
}

func (s *authService) LoginUser(user LoginRequest) (LoginResponse, error) {
	userFromDB, err := s.userRepository.FindByUsername(user.Username)

	if err != nil {
		return LoginResponse{}, errors.New("not found")
	}

	if !utils.ComparePassword(user.Password, userFromDB.Password) {
		return LoginResponse{}, errors.New("unauthorized")
	}

	accessAccessToken, accessExpiresIn, accessTokenType, err := utils.GenerateJWT(userFromDB.ID)

	if err != nil {
		return LoginResponse{}, errors.New("internal server error")
	}

	return LoginResponse{
		AccessToken: accessAccessToken,
		ExpiresIn:   accessExpiresIn,
		TokenType:   accessTokenType,
	}, nil
}
