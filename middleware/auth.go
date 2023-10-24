package Middleware

import (
	"net/http"
	ApiResponseModel "testing-api/model/api/response"

	"github.com/gin-gonic/gin"
)

func BasicAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		if authHeader == secret {
			c.Next()
			return
		}

		c.JSON(http.StatusUnauthorized, ApiResponseModel.ErrorResponse{Message: "Not authorized"})
		c.Abort()
	}
}
