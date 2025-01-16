package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mazeyqian/go-gin-gee/pkg/crypto"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("authorization")
		if !crypto.ValidateToken(authorizationHeader) {
			code := http.StatusUnauthorized
			c.AbortWithStatusJSON(code, gin.H{"code": code, "message": "unauthorized"})
			return
		} else {
			c.Next()
		}
	}
}
