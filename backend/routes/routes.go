package routes

import (
	"github.com/abhaykamath_007/library-management-system/backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	r.GET("/books", controllers.GetBooks)
	r.POST("/books/create", controllers.CreateBook)

}
