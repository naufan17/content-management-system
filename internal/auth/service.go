package auth

import (
	"errors"

	"github.com/naufan17/content-management-system/pkg/utils"
)

type AuthService interface {
	LoginUser(user LoginDto) (AccessTokenDto, error)
}

type authService struct {
	userRepository UserRepository
}

func NewAuthService(userRepository UserRepository) AuthService {
	return &authService{
		userRepository: userRepository,
	}
}

func (s *authService) LoginUser(user LoginDto) (AccessTokenDto, error) {
	userFromDB, err := s.userRepository.FindByUsername(user.Username)

	if err != nil {
		return AccessTokenDto{}, errors.New("not found")
	}

	if !utils.ComparePassword(user.Password, userFromDB.Password) {
		return AccessTokenDto{}, errors.New("unauthorized")
	}

	accessAccessToken, accessExpiresIn, accessTokenType, err := utils.GenerateJWT(userFromDB.ID)

	if err != nil {
		return AccessTokenDto{}, errors.New("internal server error")
	}

	return AccessTokenDto{
		AccessToken: accessAccessToken,
		ExpiresIn:   accessExpiresIn,
		TokenType:   accessTokenType,
	}, nil
}
