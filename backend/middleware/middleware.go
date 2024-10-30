package middleware

import (
	"net/http"
	"strings"

	"github.com/abhaykamath_007/library-management-system/backend/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the token from the Authorization header
		tokenString := c.Request.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		tokenString = strings.TrimSpace(tokenString) // Ensure no spaces

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
			c.Abort()
			return
		}

		// Parse and validate the token
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Ensure the signing method is correct
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
			}
			return []byte(config.GetSecretKey()), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
			c.Abort()
			return
		}

		// Add user info to context, accessible in next handler
		c.Set("username", claims["username"].(string))
		c.Set("user_id", claims["user_id"].(float64))
		c.Set("role", claims["role"].(string))
		c.Next()
	}
}
