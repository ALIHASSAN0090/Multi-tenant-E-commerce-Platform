package router

import (
	"ecommerce-platform/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) SignUp(c *gin.Context) {
	var req models.SignUpReq
	var signUp *models.Users

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if errMess := r.Val.ValidateReq(c, &req); len(errMess) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errMess})
		return
	}
	// if err := utils.Decode(req, &signUp); err != nil {
	// 	r.Logger.Info("failed to decode signup response: %v", err)
	// 	c.JSON(http.StatusInternalServerError, nil)
	// 	return
	// }

	if userExists, err := r.AuthService.SignUp(c, signUp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	} else if userExists {
		c.JSON(http.StatusConflict, models.Response{
			Data:    userExists,
			Message: "User already exists",
			Status:  http.StatusConflict,
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Message: "User signed up successfully",
		Status:  http.StatusOK,
	})
}
