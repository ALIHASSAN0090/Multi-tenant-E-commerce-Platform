package router

import (
	"ecommerce-platform/models"
	"ecommerce-platform/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) SignUp(c *gin.Context) {
	var req models.SignUpReq
	var signUp *models.Users

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleJsonError(c, err)
		return
	}

	if errMess := r.Val.ValidateReq(c, &req); len(errMess) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errMess})
		return
	}

	if err := utils.Decode(req, &signUp); err != nil {
		r.Logger.Info("failed to decode signup response: %v", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	if userExists, err := r.AuthService.SignUp(c, signUp); err != nil {
		utils.HandleJsonError(c, err)
		return
	} else if userExists {
		c.JSON(http.StatusConflict, models.SuccessResponse{
			Data:       userExists,
			Message:    "User already exists",
			StatusCode: http.StatusConflict,
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message:    "User signed up successfully",
		StatusCode: http.StatusOK,
	})
}

func (r *Router) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"message": "API is running smoothly",
	})
}

func (r *Router) Login(c *gin.Context) {
	var creds models.LoginReq

	if err := c.ShouldBindJSON(&creds); err != nil {
		utils.HandleJsonError(c, err)
		return
	}

	if errMess := r.Val.ValidateReq(c, &creds); len(errMess) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errMess})
		return
	}

	token, err := r.AuthService.ProcessLogin(c, &creds)
	utils.HandleJsonError(c, err)

	if err == nil {
		c.JSON(http.StatusOK, models.SuccessResponse{
			Data:       token,
			Message:    "Login successful",
			StatusCode: http.StatusOK,
		})
	} else {
		c.JSON(http.StatusUnauthorized, models.SuccessResponse{
			Message:    "Invalid credentials",
			StatusCode: http.StatusUnauthorized,
		})
	}

}
