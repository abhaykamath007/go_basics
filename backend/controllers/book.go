package controllers

import (
	"net/http"
	"strconv"

	"github.com/abhaykamath_007/library-management-system/backend/models"
	"github.com/abhaykamath_007/library-management-system/backend/service"
	"github.com/abhaykamath_007/library-management-system/backend/validators"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {

	books, err := service.GetBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books", "details": err.Error()})
		return
	}

	if len(books) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No books found"})
	} else {
		c.JSON(http.StatusOK, books)
	}
}

func GetBookByID(c *gin.Context) {
	idParam := c.Param("id")
	bookID, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	book, err := service.GetBookByID(bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve book", "details": err.Error()})
		return
	}

	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
	} else {
		c.JSON(http.StatusOK, book)
	}

}

func CreateBook(c *gin.Context) {
	var newBook models.Book

	// Bind JSON input
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	if err := validators.ValidateBook(newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Proceed to create the book
	createdBook, err := service.CreateBook(newBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdBook)
}

func UpdateBook(c *gin.Context) {
	// Parse book ID from URL parameter
	bookID := c.Param("id")
	id, err := strconv.Atoi(bookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var updatedBook models.Book
	// Bind JSON input to `updatedBook`
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	// Set the parsed ID to `updatedBook`
	updatedBook.ID = id

	// Validate fields of `updatedBook`
	if err := validators.ValidateBook(updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Proceed to update the book
	updatedBook, err = service.UpdateBook(updatedBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully", "book": updatedBook})
}
