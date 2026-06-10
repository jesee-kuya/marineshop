package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jesee-kuya/marineshop/domain"
)

func GenerateToken(user *domain.User, cfg *domain.JWTConfig) (string, error) {
	claims := &domain.Claims{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(cfg.TokenExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.SecretKey))
}
