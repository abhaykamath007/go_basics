package routes

import (
	"restapi/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/albums", controllers.GetAlbums)
	router.GET("/albums/:id", controllers.GetAlbumByID)
	router.POST("/albums", controllers.PostAlbum)
	router.PUT("/albums/:id", controllers.UpdateAlbumByID)
	router.DELETE("/albums/:id", controllers.DeleteAlbumByID)
}
