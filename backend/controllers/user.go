package controllers

import (
	"net/http"
	"strconv"

	"github.com/abhaykamath_007/library-management-system/backend/service"
	"github.com/gin-gonic/gin"
)

func GetBorrowedBooks(c *gin.Context) {
	userIDParam := c.Param("user_id")
	userID, err := strconv.Atoi(userIDParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	borrowedBooks, err := service.GetBorrowedBooks(userID)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "user does not exist"})

		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch borrowed books", "details": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, borrowedBooks)
}

func GetReturnedBooks(c *gin.Context) {
	userIDParam := c.Param("user_id")
	userID, err := strconv.Atoi(userIDParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	returnedBooks, err := service.GetReturnedBooks(userID)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "user does not exist"})

		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch borrowed books", "details": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, returnedBooks)
}

func GetStats(c *gin.Context) {
	stats, err := service.GetStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch stats", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stats)
}
