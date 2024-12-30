package middleware

import (
	"ecommerce-platform/Dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StatusCheck(dao Dao.SellerDao) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIdInterface, exists := c.Get("Id")
		if !exists {
			abortWithError(c, http.StatusUnauthorized, "Access denied", "Unauthorized")
			return
		}

		userId, ok := userIdInterface.(int64)
		if !ok {
			abortWithError(c, http.StatusInternalServerError, "Internal server error", "User ID is not of type int64")
			return
		}

		user, err := dao.IsActive(c, userId)
		if err != nil || !user {
			abortWithError(c, http.StatusUnauthorized, "Access denied", "Your Account status is disable")
			return
		}

		c.Next()
	}
}

func abortWithError(c *gin.Context, statusCode int, message, detail string) {
	c.AbortWithStatusJSON(statusCode, gin.H{
		"error": gin.H{
			"statusCode": statusCode,
			"message":    message,
			"detail":     detail,
		},
	})
}
