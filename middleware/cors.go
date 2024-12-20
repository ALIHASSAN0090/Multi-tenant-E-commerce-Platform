package middleware

import "github.com/gin-gonic/gin"

func EnableCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedOrigins := []string{"*"}
		origin := c.Request.Header.Get("Origin")

		allowOrigin := ""
		for _, o := range allowedOrigins {
			if origin == o {
				allowOrigin = origin
				break
			}
		}

		if allowOrigin == "" {
			allowOrigin = "*"
		}

		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
