package middleware

import (
	"net/http"
	"strings"

	"jna-manager/internal/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "User Not Authorized to perform action",
			})
			return
		}

		// Check if the token starts with "Bearer " and extract the token
		const bearerPrefix = "Bearer "
		if !strings.HasPrefix(tokenString, bearerPrefix) {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token format",
			})
			return
		}

		tokenString = strings.TrimPrefix(tokenString, bearerPrefix)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return cfg.SecretKey, nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "User Not Authorized to perform action",
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "User Not Authorized to perform action",
			})
			return
		}

		userRole, ok := claims["user_role"].(string)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "User Not Authorized to perform action",
			})
			return
		}

		// Check the user's role against the required roles for the current endpoint
		requiredRoles := c.GetStringSlice("requiredRoles")
		if len(requiredRoles) > 0 && !containsRole(requiredRoles, userRole) {
			c.AbortWithStatus(http.StatusForbidden)
			c.JSON(http.StatusForbidden, gin.H{
				"message": "User is forbidden to perform action",
			})
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}

func RequireRoles(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("requiredRoles", roles)
		c.Next()
	}
}

func containsRole(roles []string, role string) bool {
	for _, r := range roles {
		if strings.EqualFold(r, role) {
			return true
		}
	}
	return false
}
