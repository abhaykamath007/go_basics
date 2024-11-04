package service

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/abhaykamath_007/library-management-system/backend/models"
	"github.com/abhaykamath_007/library-management-system/backend/repo"
)

func BorrowBook(userID, bookID int) error {
	book, err := repo.FetchBookByID(bookID)
	log.Print("book error : ", err)
	if err != nil {
		return fmt.Errorf("failed to fetch book")
	}
	if book == nil {
		return fmt.Errorf("invalid book id")
	}

	if book.AvailabilityStatus != "available" {
		return errors.New("book is currently unavailable")
	}

	dueDate := time.Now().AddDate(0, 0, 14)

	if err := repo.CreateTransaction(userID, bookID, dueDate); err != nil {
		return fmt.Errorf("failed to create transaction: %w", err)
	}

	return repo.UpdateBookAvailability(bookID, "checked_out")
}

func ReturnBook(userID, bookID int) error {
	book, err := repo.FetchBookByID(bookID)
	if err != nil {
		return fmt.Errorf("failed to fetch book: %w", err)
	}
	if book == nil {
		return errors.New("book does not exist")
	}

	transaction := &models.Transaction{}
	err = repo.FetchTransactionByUserAndBook(transaction, userID, bookID)
	if err != nil {
		return fmt.Errorf("failed to fetch transaction: %w", err)
	}
	if transaction.ID == 0 {
		return errors.New("book is not currently borrowed by this user")
	}

	if err := repo.MarkTransactionAsReturned(transaction.ID); err != nil {
		return fmt.Errorf("failed to mark transaction as returned: %w", err)
	}

	return repo.UpdateBookAvailability(bookID, "available")
}
