package middleware

import (
	"ecommerce-platform/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AdminTokenHeaderKey          = "Authorization"
	AdminAuthorizationTypeBearer = "BEARER"
)

func Auth(allowedRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(AdminTokenHeaderKey)
		if len(token) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Token not found in header"})
			return
		}

		tokenStr := strings.TrimPrefix(token, "Bearer ")
		claims, err := ValidateAccessToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Error: err.Error()})
			return
		}

		if !AllowedRoles(claims.Role, allowedRoles) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Unauthorized access"})
			return
		}

		setClaims(c, *claims)
		c.Next()
	}
}

func AllowedRoles(role *string, allowedRoles []string) bool {
	if role == nil {
		return false
	}
	for _, r := range allowedRoles {
		if r == *role {
			return true
		}
	}
	return false
}

func setClaims(c *gin.Context, claims models.CustomClaims) {
	c.Set("Id", claims.Id)
	c.Set("email", claims.Email)
	c.Set("role", claims.Role)
}
