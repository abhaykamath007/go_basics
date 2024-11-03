package repo

import (
	"github.com/abhaykamath_007/library-management-system/backend/database"
	"github.com/abhaykamath_007/library-management-system/backend/models"
)

// CheckUsernameExists checks if a username is already in use
func CheckUsernameExists(username string) (bool, error) {
	var existingUser models.User
	sqlQuery := "SELECT username FROM users WHERE username = ?"
	err := database.DB.Raw(sqlQuery, username).Scan(&existingUser).Error

	if err != nil {
		return false, err
	}

	return existingUser.Username != "", nil
}

// CreateUser inserts a new user into the database
func CreateUser(newUser models.User) error {
	sql := "INSERT INTO users (username, password, role) VALUES (?, ?, ?)"
	return database.DB.Exec(sql, newUser.Username, newUser.Password, "Member").Error
}

func FindUserByUsername(username string) (models.User, error) {
	var user models.User
	sql := "SELECT id, username, password, role FROM users WHERE username = ?"
	err := database.DB.Raw(sql, username).Scan(&user).Error
	return user, err
}
