package service

import (
	"errors"
	"fmt"

	"github.com/abhaykamath_007/library-management-system/backend/models"
	"github.com/abhaykamath_007/library-management-system/backend/repo"
	"github.com/abhaykamath_007/library-management-system/backend/utils"
)

func RegisterUser(newUser models.User) error {
	// Check if the username already exists
	exists, err := repo.CheckUsernameExists(newUser.Username)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("username already exists")
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		return err
	}
	newUser.Password = hashedPassword

	// Call the repository to save the user
	return repo.CreateUser(newUser)
}

func LoginUser(username, password string) (string, int64, int, error) {
	user, err := repo.FindUserByUsername(username)
	if err != nil {
		return "", 0, 0, errors.New("user not found")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", 0, 0, errors.New("invalid credentials")
	}

	token, expTime, err := utils.GenerateJWT(user)
	if err != nil {
		return "", 0, 0, errors.New("failed to generate token")
	}

	return token, expTime, user.ID, nil
}
