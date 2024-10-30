package utils

import (
	"time"

	"github.com/abhaykamath_007/library-management-system/backend/config"
	"github.com/abhaykamath_007/library-management-system/backend/models"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(user models.User) (string, int64, error) {
	expirationTime := time.Now().Add(time.Hour * 72).Unix() // 72-hour expiration time

	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      expirationTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.GetSecretKey()))
	return tokenString, expirationTime, err
}
