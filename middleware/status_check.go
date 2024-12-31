package middleware

import (
	"ecommerce-platform/Dao"
	"ecommerce-platform/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Db *Dao.SellerDaoImpl

func StatusCheck(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {

		userId, err := utils.GetContextId(c)
		utils.HandleError(err)

		if Db == nil {

			abortWithError(c, http.StatusInternalServerError, "Internal server error", "DAO is not initialized")
			return
		}

		user, err := Db.IsActive(c, userId)
		if err != nil || !user {
			abortWithError(c, http.StatusUnauthorized, "Access denied", "Your Account status is disabled")
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
