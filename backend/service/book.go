package service

import (
	"errors"

	"github.com/abhaykamath_007/library-management-system/backend/models"
	"github.com/abhaykamath_007/library-management-system/backend/repo"
)

func GetBookByID(bookID int) (*models.Book, error) {
	return repo.FetchBookByID(bookID)
}

func GetBooks() ([]models.Book, error) {
	return repo.FetchBooks()
}

// CreateBook inserts a new book into the database.
func CreateBook(newBook models.Book) (models.Book, error) {
	return repo.InsertBook(newBook)
}

// UpdateBook updates a book's information.
func UpdateBook(updatedBook models.Book) (models.Book, error) {
	// Verify that the book exists
	if exists, err := repo.BookExistsByID(updatedBook.ID); err != nil {
		return models.Book{}, err
	} else if !exists {
		return models.Book{}, errors.New("book not found")
	}

	// Update the book in the repository
	return repo.UpdateBook(updatedBook)
}
