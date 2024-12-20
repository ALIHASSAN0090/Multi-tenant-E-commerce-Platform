package middleware

import (
	configs "ecommerce-platform/configs"
	"ecommerce-platform/models"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte(configs.AppConfig.JWT_SECRET)

const (
	TokenHeaderKey          = "Authorization"
	AuthorizationTypeBearer = "BEARER"
	accessTokenDuration     = 24 * time.Hour
)

func GenerateAccessToken(user *models.Users) (string, error) {
	claims := models.CustomClaims{
		Id:    uint(user.ID),
		Email: &user.Email,
		Name:  &user.UserName,
		Role:  &user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessTokenDuration).Unix(),
		},
	}
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaim.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}
	return token, nil
}

func ValidateAccessToken(tokenString string) (*models.CustomClaims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method in access token")
		}
		return []byte(jwtSecret), nil
	}

	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, keyFunc)
	if err != nil {
		return nil, fmt.Errorf("error parsing access token claims: %w", err)
	}

	claims, ok := token.Claims.(*models.CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid access token")
	}

	return claims, nil
}
