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
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{
				Error: models.Error{
					StatusCode: http.StatusUnauthorized,
					Message:    "Token not found in header",
					Detail:     "The request does not contain a valid authorization token.",
				},
			})
			return
		}

		tokenStr := strings.TrimPrefix(token, "Bearer ")
		claims, err := ValidateAccessToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{
				Error: models.Error{
					StatusCode: http.StatusUnauthorized,
					Message:    "Invalid token",
					Detail:     err.Error(),
				},
			})
			return
		}

		if !AllowedRoles(claims.Role, allowedRoles) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{
				Error: models.Error{
					StatusCode: http.StatusUnauthorized,
					Message:    "Unauthorized access",
					Detail:     "The user's role does not have permission to access this resource.",
				},
			})
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
