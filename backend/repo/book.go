package repo

import (
	"github.com/abhaykamath_007/library-management-system/backend/database"
	"github.com/abhaykamath_007/library-management-system/backend/models"
)

func FetchBookByID(bookID int) (*models.Book, error) {
	var book models.Book
	sql := "SELECT * FROM books WHERE id = ?"

	if err := database.DB.Raw(sql, bookID).Scan(&book).Error; err != nil {
		return nil, err
	}

	if book.ID == 0 {
		return nil, nil
	}

	return &book, nil
}

func FetchBooks() ([]models.Book, error) {
	var books []models.Book
	sql := "SELECT * FROM books"

	if err := database.DB.Raw(sql).Scan(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func BookExists(title, author string) (bool, error) {
	var existingBook models.Book
	sqlQuery := "SELECT * FROM books WHERE title = ? AND author = ? LIMIT 1"
	if err := database.DB.Raw(sqlQuery, title, author).Scan(&existingBook).Error; err != nil {
		return false, err
	}
	return existingBook.ID != 0, nil
}

func InsertBook(newBook models.Book) (models.Book, error) {

	newBook.AvailabilityStatus = "Available"
	sql := "INSERT INTO books (title, author, genre, publication_year, availability_status) VALUES (?, ?, ?, ?, ?)"
	if err := database.DB.Exec(sql, newBook.Title, newBook.Author, newBook.Genre, newBook.PublicationYear, newBook.AvailabilityStatus).Error; err != nil {
		return models.Book{}, err
	}

	var createdBook models.Book
	database.DB.Raw("SELECT * FROM books WHERE id = LAST_INSERT_ID()").Scan(&createdBook)
	return createdBook, nil
}

// BookExistsByID checks if a book exists by its ID.
func BookExistsByID(id int) (bool, error) {
	var count int64
	sqlCheck := "SELECT COUNT(*) FROM books WHERE id = ?"
	if err := database.DB.Raw(sqlCheck, id).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// UpdateBook performs the actual update on the book record.
func UpdateBook(updatedBook models.Book) (models.Book, error) {
	sqlUpdate := `
		UPDATE books
		SET title = ?, author = ?, genre = ?, publication_year = ?, availability_status = ?
		WHERE id = ?
	`
	if err := database.DB.Exec(sqlUpdate, updatedBook.Title, updatedBook.Author, updatedBook.Genre, updatedBook.PublicationYear, updatedBook.AvailabilityStatus, updatedBook.ID).Error; err != nil {
		return models.Book{}, err
	}

	return updatedBook, nil
}
