package validators

import (
	"errors"
	"time"

	"github.com/abhaykamath_007/library-management-system/backend/models"
)

func ValidateBook(book models.Book) error {
	if book.Title == "" {
		return errors.New("title is required")
	}
	if book.Author == "" {
		return errors.New("author is required")
	}
	if book.Genre == "" {
		return errors.New("genre is required")
	}
	if book.PublicationYear == 0 {
		return errors.New("publication year is required")
	}
	currentYear := time.Now().Year()
	if book.PublicationYear > currentYear {
		return errors.New("publication year cannot be in the future")
	}
	return nil
}
