package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/api/v1/auth/login" {
			c.Next()
		} else {
			h := AuthHeader{}
			if err := c.ShouldBindHeader(&h); err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": err.Error(),
				})
				c.Abort()
			}

			if h.AuthorizationHeader == "ini_adalah_token_yang_valid" {
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "token invalid",
				})
				c.Abort()
			}
		}
	}
}
