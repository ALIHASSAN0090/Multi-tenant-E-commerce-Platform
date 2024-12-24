package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ecommerce-platform/router"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPublicHealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)

	engine := gin.New()
	r := &router.Router{
		Engine: engine,
	}

	r.SetupRoutes()

	req, _ := http.NewRequest("GET", "/public/health-check", nil)

	w := httptest.NewRecorder()

	r.Engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	expectedResponse := map[string]string{
		"message": "API is running smoothly",
		"status":  "healthy",
	}

	assert.Equal(t, expectedResponse, response)
}
