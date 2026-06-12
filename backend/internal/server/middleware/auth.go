package middleware

import (
	//"net/http"
	"strings"

	"helloblog/internal/pkg/jwt"
	"helloblog/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func Auth(jwtManager *jwt.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Fail(c, response.NewError(response.CodeUnauthorized, "missing authorization header"))
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			response.Fail(c, response.NewError(response.CodeUnauthorized, "invalid authorization header"))
			c.Abort()
			return
		}

		claims, err := jwtManager.Validate(parts[1])
		if err != nil {
			response.Fail(c, response.NewError(response.CodeUnauthorized, "invalid or expired token"))
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
