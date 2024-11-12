package repo

import (
	"github.com/abhaykamath_007/library-management-system/backend/database"
	"github.com/abhaykamath_007/library-management-system/backend/models"
)

func CheckUserExists(userID int) (bool, error) {
	var exists bool
	sql := "SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)"
	if err := database.DB.Raw(sql, userID).Scan(&exists).Error; err != nil {
		return false, err
	}
	return exists, nil
}

func FetchBorrowedBooks(userID int) ([]models.BorrowedBook, error) {
	var books []models.BorrowedBook

	sql := `SELECT b.id,b.title,b.author,t.borrowed_date,t.due_date
	        FROM books AS b
			JOIN transactions AS t ON b.id = t.book_id
			WHERE t.user_id = ? AND t.returned = false`

	if err := database.DB.Raw(sql, userID).Scan(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func FetchReturnedBooks(userID int) ([]models.BorrowedBook, error) {
	var books []models.BorrowedBook

	sql := `SELECT b.id,b.title,b.author,t.borrowed_date,t.due_date
	        FROM books AS b
			JOIN transactions AS t ON b.id = t.book_id
			WHERE t.user_id = ? AND t.returned = true`

	if err := database.DB.Raw(sql, userID).Scan(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
