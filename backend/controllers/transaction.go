package controllers

import (
	"net/http"
	"strconv"

	"github.com/abhaykamath_007/library-management-system/backend/service"
	"github.com/gin-gonic/gin"
)

func BorrowBook(c *gin.Context) {

	var request struct {
		UserID int `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	bookIDParam := c.Param("book_id")
	bookID, err := strconv.Atoi(bookIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	if err := service.BorrowBook(request.UserID, bookID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to borrow book", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book borrowed successfully"})
}

func ReturnBook(c *gin.Context) {
	var request struct {
		UserID int `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}
	bookIDParam := c.Param("book_id")
	bookID, err := strconv.Atoi(bookIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	if err := service.ReturnBook(request.UserID, bookID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to return book", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book returned successfully"})
}
