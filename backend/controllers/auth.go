package controllers

import (
	"net/http"

	"github.com/abhaykamath_007/library-management-system/backend/models"
	"github.com/abhaykamath_007/library-management-system/backend/service"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	// Call the service layer to register the user
	err := service.RegisterUser(newUser)
	if err != nil {
		if err.Error() == "username already exists" {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "details": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var request models.LoginRequest // Using LoginRequest struct for binding

	// Bind JSON input
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	// Call the service to handle login
	token, expTime, err := service.LoginUser(request.Username, request.Password) // Call to service layer
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}
		if err.Error() == "invalid credentials" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login", "details": err.Error()})
		return
	}

	// Respond with the token
	c.JSON(http.StatusOK, gin.H{
		"token":          token,
		"expirationTime": expTime,
	})
}
