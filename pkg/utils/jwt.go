package utils

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/naufan17/content-management-system/pkg/config"
)

var (
	jwtAccessSecret = []byte(config.LoadConfig().JWTAccessSecret)
	jwtAccessExpStr = config.LoadConfig().JWTAccessExp
	jwtAccessExp, _ = strconv.Atoi(jwtAccessExpStr)
)

type Claims struct {
	Sub uuid.UUID `json:"sub"`
	Iat int64     `json:"iat"`
	jwt.StandardClaims
}

func GenerateJWT(id uuid.UUID) (string, int64, string, error) {
	expirationTime := time.Now().Add(time.Duration(jwtAccessExp) * time.Millisecond)

	claims := &Claims{
		Sub: id,
		Iat: time.Now().Unix(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtAccessSecret)

	if err != nil {
		return "", 0, "", err
	}

	return tokenString, expirationTime.Unix(), "Bearer", nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtAccessSecret, nil
	})

	if claims, ok := token.Claims.(*Claims); !ok && !token.Valid {
		return claims, nil
	}

	return nil, err
}
