package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if header == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Authorization header required", "error": true})
			ctx.Abort()
			return
		}

		if !strings.HasPrefix(header, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "invalid header type", "error": true})
			ctx.Abort()
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")

		checkToken, err := VerifyToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error(), "error": true})
			ctx.Abort()
			return
		}

		ctx.Set("id", checkToken.UserId)
		ctx.Set("role", checkToken.Role)

		ctx.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRole, exists := ctx.Get("role")
		fmt.Println("Role from Token:", userRole)

		if !exists || userRole != "admin" {
			ctx.JSON(http.StatusForbidden, gin.H{"status": 403, "message": "Forbidden: Admins only", "error": true})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
