package repo

import (
	"time"

	"github.com/abhaykamath_007/library-management-system/backend/database"
	"github.com/abhaykamath_007/library-management-system/backend/models"
)

func CreateTransaction(userID, bookID int, dueDate time.Time) error {
	borrowedDate := time.Now()
	sql := "INSERT INTO transactions (user_id, book_id, borrowed_date, due_date, returned) VALUES (?, ?, ?, ?, 0)"
	return database.DB.Exec(sql, userID, bookID, borrowedDate, dueDate).Error
}

func UpdateBookAvailability(bookID int, status string) error {
	sql := "UPDATE books SET availability_status = ? WHERE id = ?"
	return database.DB.Exec(sql, status, bookID).Error
}

func MarkTransactionAsReturned(transactionID int) error {
	sql := "UPDATE transactions SET returned = 1 WHERE id = ?"
	return database.DB.Exec(sql, transactionID).Error
}

func FetchTransactionByUserAndBook(transaction *models.Transaction, userID, bookID int) error {
	sql := "SELECT * FROM transactions WHERE user_id = ? AND book_id = ? AND returned = 0"
	err := database.DB.Raw(sql, userID, bookID).Scan(transaction).Error
	if err != nil {
		return err
	}
	return nil
}
