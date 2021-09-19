package routes

import (
	"github.com/Rahmanwghazi/AlterraCRUD/controllers"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/books/", controllers.GetBookController)
	e.GET("/books/:id", controllers.GetBookByIdController)
	e.POST("/books/", controllers.AddBookController)
	e.DELETE("/books/:id", controllers.DeleteBookByIdController)
	e.PUT("/books/:id", controllers.UpdateBookByIdController)
	return e
}
