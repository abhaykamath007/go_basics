package controllers

import (
	"net/http"

	"github.com/abhaykamath_007/library-management-system/backend/database"
	"github.com/abhaykamath_007/library-management-system/backend/models"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {

	var books []models.Book

	sql := "SELECT * FROM books"

	if err := database.DB.Raw(sql).Scan(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books", "details": err.Error()})
		return
	}

	if len(books) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No books found"})
	} else {
		c.JSON(http.StatusOK, books)
	}
}

func CreateBook(c *gin.Context) {
	var newBook models.Book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	sql := "INSERT INTO books (title, author, genre, publication_year) VALUES (?, ?, ?, ?)"

	if err := database.DB.Exec(sql, newBook.Title, newBook.Author, newBook.Genre, newBook.PublicationYear).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newBook)

}
