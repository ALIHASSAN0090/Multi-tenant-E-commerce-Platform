package models

import (
	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	Id    uint    `json:"id"`
	Email *string `json:"email"`
	Name  *string `json:"name"`
	Role  *string `json:"role"`
	jwt.StandardClaims
}
