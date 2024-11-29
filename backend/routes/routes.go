package routes

import (
	"github.com/abhaykamath_007/library-management-system/backend/controllers"
	"github.com/abhaykamath_007/library-management-system/backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	authorized := r.Group("/").Use(middleware.AuthMiddleware())
	{
		authorized.GET("/books", controllers.GetBooks)
		authorized.GET("/books/:id", controllers.GetBookByID)
		authorized.POST("/books/:book_id/borrow", controllers.BorrowBook)
		authorized.POST("/books/:book_id/return", controllers.ReturnBook)

		authorized.GET("/genres", controllers.GetGenres)

		authorized.GET("users/books/:user_id/borrowed", controllers.GetBorrowedBooks)
		authorized.GET("users/books/:user_id/returned", controllers.GetReturnedBooks)

	}

	admin := r.Group("/admin").Use(middleware.AdminOnlyMiddleware())
	{
		admin.POST("/books/create", controllers.CreateBook)
		admin.PUT("books/update/:id", controllers.UpdateBook)
	}

}
