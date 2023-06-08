package controllers

import (
	"go-latihan/database"
	"go-latihan/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all Books
func GetBooksController(c echo.Context) error {
	books, e := database.GetBooks()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	BookId := GetBookIdFromParam(c)
	book, e := database.GetBookById(BookId)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book by id",
		"book":    book,
	})
}

// create new book
func CreateBookController(c echo.Context) error {
	book := &models.Book{}

	if err := c.Bind(book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	createdBook, err := database.Createbook(book)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    createdBook,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	BookId := GetBookIdFromParam(c)
	book, e := database.DeleteBookById(BookId)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Delete book",
		"book":    book,
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	BookId := GetIdFromParam(c)
	updatedBook := &models.Book{}

	if e := c.Bind(updatedBook); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	book, e := database.UpdateBookById(BookId, updatedBook)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"book":    book,
	})
}

func GetBookIdFromParam(c echo.Context) int {
	BookId, _ := strconv.Atoi(c.Param("id"))
	return BookId
}
