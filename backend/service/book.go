package service

import (
	"github.com/abhaykamath_007/library-management-system/backend/models"
	"github.com/abhaykamath_007/library-management-system/backend/repo"
)

func GetBookByID(bookID int) (*models.Book, error) {
	return repo.FetchBookByID(bookID)
}

func GetBooks() ([]models.Book, error) {
	return repo.FetchBooks()
}

func CheckBookExists(title, author string) (bool, error) {
	return repo.BookExists(title, author)
}

// CreateBook inserts a new book into the database.
func CreateBook(newBook models.Book) (models.Book, error) {
	return repo.InsertBook(newBook)
}
