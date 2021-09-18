package controllers

import (
	"net/http"
	"strconv"

	"github.com/Rahmanwghazi/AlterraCRUD/config"
	"github.com/Rahmanwghazi/AlterraCRUD/models"
	"github.com/labstack/echo/v4"
)

func GetBookController(e echo.Context) error {
	var books []models.Book
	err := config.InitDB().Find(&books).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "successful operation",
		"data":    books,
	})
}

func AddBookController(c echo.Context) error {

	book := models.Book{}
	c.Bind(&book)

	err := config.InitDB().Save(&book).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successful operation",
		"book":    book,
	})
}

func DeleteBookByIdController(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	book := &models.Book{}

	err := config.InitDB().Find(&book, "id = ?", id).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	config.InitDB().Delete(&models.Book{}, id)

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "successful deleted",
	})
}
