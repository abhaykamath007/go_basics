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

func CountUsers() (int64, error) {
	var count int64

	sql := `SELECT count(*) FROM users WHERE role != 'admin' `
	if err := database.DB.Raw(sql).Scan(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func CountBooks() (int64, error) {
	var count int64

	sql := `SELECT count(*) FROM books`
	if err := database.DB.Raw(sql).Scan(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func CountAvailableBooks() (int64, error) {
	var count int64

	sql := `SELECT count(*) FROM books where availability_status='available'`
	if err := database.DB.Raw(sql).Scan(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func CountBorrowedBooks() (int64, error) {
	var count int64

	sql := `SELECT count(*) FROM books where availability_status != 'available'`
	if err := database.DB.Raw(sql).Scan(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func CountBooksBorrowedThisMonth() (int64, error) {
	var count int64

	sql := `SELECT COUNT(*) FROM 
	transactions WHERE MONTH(borrowed_date) = MONTH(CURRENT_DATE())
	AND YEAR(borrowed_date) = YEAR(CURRENT_DATE())`

	err := database.DB.Raw(sql).Scan(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CountBooksAddedThisMonth() (int64, error) {
	var count int64

	sql := `SELECT COUNT(*) FROM books WHERE MONTH(created_at) = MONTH(CURRENT_DATE())
	AND YEAR(created_at) = YEAR(CURRENT_DATE())`

	err := database.DB.Raw(sql).Scan(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
