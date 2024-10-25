package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetDB(database *gorm.DB) {
	db = database
}

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func GetAlbums(c *gin.Context) {
	var albums []Album
	if err := db.Raw("select * from albums").Scan(&albums).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve albums"})
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbum(c *gin.Context) {
	var newAlbum Album
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := db.Exec("insert into albums (id, title, artist, price) values (?, ?, ?, ?)", newAlbum.ID, newAlbum.Title, newAlbum.Artist, newAlbum.Price).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create album"})
		return
	}
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var album Album
	if err := db.Raw("select id, title, artist, price from albums where id = ?", id).Scan(&album).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func UpdateAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum Album

	if err := c.BindJSON(&updatedAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid input"})
		return
	}

	if err := db.Exec("update albums set title = ?, artist = ?, price = ? where id = ?", updatedAlbum.Title, updatedAlbum.Artist, updatedAlbum.Price, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found or update failed"})
		return
	}
	c.IndentedJSON(http.StatusOK, updatedAlbum)
}

func DeleteAlbumByID(c *gin.Context) {
	id := c.Param("id")
	if err := db.Exec("delete from albums where id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
}
