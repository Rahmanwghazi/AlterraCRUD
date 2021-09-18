package controllers

import (
	"net/http"
	"strconv"

	"github.com/Rahmanwghazi/AlterraCRUD/models"
	"github.com/labstack/echo/v4"
)

func GetBookController(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	book := models.Book{ID: id, Title: "Buku Baru", ISBN: "123213", Writer: "Andy"}
	return e.JSON(http.StatusOK, book)
}

func AddBookController(c echo.Context) error {
	/* id := c.FormValue("id")
	title := c.FormValue("title")
	isbn := c.FormValue("isbn")
	writer := c.FormValue("writer") */

	book := models.Book{}
	c.Bind(&book)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"book": book,
	})
}
