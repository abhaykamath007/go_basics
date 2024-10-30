package service

import (
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
