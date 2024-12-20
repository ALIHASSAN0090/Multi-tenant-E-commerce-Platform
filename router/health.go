package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Service is up"})
}
