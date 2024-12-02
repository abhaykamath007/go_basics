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

func GetStats() (map[string]int64, error) {
	stats := make(map[string]int64)

	totalUsers, err := repo.CountUsers()
	if err != nil {
		return nil, err
	}
	stats["total_users"] = totalUsers

	totalBooks, err := repo.CountBooks()
	if err != nil {
		return nil, err
	}
	stats["total_books"] = totalBooks

	borrowedBooks, err := repo.CountBorrowedBooks()
	if err != nil {
		return nil, err
	}
	stats["borrowed_books"] = borrowedBooks

	availableBooks, err := repo.CountAvailableBooks()
	if err != nil {
		return nil, err
	}
	stats["available_books"] = availableBooks

	booksBorrowedThisMonth, err := repo.CountBooksBorrowedThisMonth()
	if err != nil {
		return nil, err
	}
	stats["books_borrowed_this_month"] = booksBorrowedThisMonth

	booksAddedThisMonth, err := repo.CountBooksAddedThisMonth()
	if err != nil {
		return nil, err
	}
	stats["books_added_this_month"] = booksAddedThisMonth

	return stats, nil
}
