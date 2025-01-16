package router

import (
	"context"
	config "ecommerce-platform/configs"
	"ecommerce-platform/models"
	"ecommerce-platform/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthStateString = config.Cfg.JWTSecret
)

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8005/public/callback",
	ClientID:     config.Cfg.GoogleClientID,
	ClientSecret: config.Cfg.GoogleClientSecret,
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

func (r *Router) GoogleLogin(c *gin.Context) {

	url := googleOauthConfig.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)

	utils.OpenURLInBrowser(url)

}

func (r *Router) HandleoauthCallback(c *gin.Context) {

	if c.Query("state") != oauthStateString {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	token, err := googleOauthConfig.Exchange(context.Background(), c.Query("code"))
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	client := googleOauthConfig.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	defer response.Body.Close()

	var userInfo models.OauthUserInfo
	err = json.NewDecoder(response.Body).Decode(&userInfo)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	valid := r.Val.ValidateOauthCreds(userInfo)
	if !valid {
		return
	}

	Jwttoken, err := r.AuthService.OauthSetup(c, &userInfo)
	if err == nil {
		c.JSON(http.StatusOK, models.TokenResponse{
			Token:      Jwttoken,
			Message:    "Login successful",
			StatusCode: http.StatusOK,
		})
	} else {
		c.JSON(http.StatusUnauthorized, models.TokenResponse{
			Message:    "Invalid credentials",
			StatusCode: http.StatusUnauthorized,
		})
	}
}

func (r *Router) SignUp(c *gin.Context) {
	var req models.SignUpReq
	var signUp *models.Users

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleJsonError(c, err)
		return
	}
	req.Role = "user"

	if errMess := r.Val.ValidateReq(c, &req); len(errMess) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errMess})
		return
	}

	if err := utils.Decode(req, &signUp); err != nil {
		r.Logger.Info("failed to decode signup response: %v", err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	user, message, err := r.AuthService.SignUp(c, signUp)
	utils.HandleError(err)

	c.JSON(http.StatusOK, models.SuccessResponse{
		Data:       user,
		Message:    message,
		SubMessage: message,
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
		c.JSON(http.StatusOK, models.TokenResponse{
			Token:      token,
			Message:    "Login successful",
			StatusCode: http.StatusOK,
		})
	} else {
		c.JSON(http.StatusUnauthorized, models.TokenResponse{
			Message:    "Invalid credentials",
			StatusCode: http.StatusUnauthorized,
		})
	}

}

func GetContextID(c *gin.Context) (int64, bool) {
	IDInterface, exists := c.Get("Id")
	fmt.Printf("found id :%v, type: %T\n", IDInterface, IDInterface)
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID not found"})
		return 0, false
	}
	if IDUint, ok := IDInterface.(uint); ok {
		ID := int64(IDUint)

		return ID, true
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "ID is not valid"})
	return 0, false
}
