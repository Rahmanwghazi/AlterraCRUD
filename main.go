package main

import (
	"github.com/Rahmanwghazi/AlterraCRUD/controllers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/books/:id", controllers.GetBookController)
	e.POST("/books/", controllers.AddBookController)
	e.Start(":8000")
}
