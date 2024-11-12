package service

import (
	"fmt"

	"github.com/abhaykamath_007/library-management-system/backend/models"
	"github.com/abhaykamath_007/library-management-system/backend/repo"
)

func GetBorrowedBooks(userID int) ([]models.BorrowedBook, error) {

	exists, err := repo.CheckUserExists(userID)
	if err != nil {
		return nil, fmt.Errorf("error checking user existence : %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("user not found")
	}

	borrowedBooks, err := repo.FetchBorrowedBooks(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch borrowed books: %w", err)
	}
	return borrowedBooks, nil
}

func GetReturnedBooks(userID int) ([]models.BorrowedBook, error) {

	exists, err := repo.CheckUserExists(userID)
	if err != nil {
		return nil, fmt.Errorf("error checking user existence : %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("user not found")
	}

	borrowedBooks, err := repo.FetchReturnedBooks(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch borrowed books: %w", err)
	}
	return borrowedBooks, nil
}
