package controllers

import (
	"net/http"
	"strconv"

	"github.com/Rahmanwghazi/AlterraCRUD/config"
	"github.com/Rahmanwghazi/AlterraCRUD/models"
	"github.com/labstack/echo/v4"
)

func GetBookController(c echo.Context) error {
	var books []models.Book
	err := config.InitDB().Find(&books).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": books,
	})
}

func GetBookByIdController(c echo.Context) error {
	var book models.Book

	id, _ := strconv.Atoi(c.Param("id"))

	err := config.InitDB().Table("books").First(&book, id).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"book": book,
	})
}

func AddBookController(c echo.Context) error {

	book := models.Book{}
	c.Bind(&book)

	err := config.InitDB().Save(&book).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"book": book,
	})
}

func DeleteBookByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book := models.Book{}

	err := config.InitDB().Find(&book, "id = ?", id).Error
	if err != nil {
		return c.JSON(http.StatusNoContent, map[string]interface{}{
			"message": err.Error(),
		})
	}
	config.InitDB().Delete(&models.Book{}, id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "deleted",
	})
}

func UpdateBookByIdController(c echo.Context) error {
	var book models.Book
	var books models.Book
	c.Bind(&book)

	id, _ := strconv.Atoi(c.Param("id"))

	err := config.InitDB().Table("books").First(&books, id).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err.Error(),
		})
	}

	err2 := config.InitDB().Table("books").Where("id = ?", id).Updates(book).Error

	if err2 != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"book": book,
	})
}
