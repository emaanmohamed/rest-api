package handlers

import (
	"errors"
	"github.com/emaanmohamed/rest-api/dto"
	"github.com/labstack/echo/v4"
	"strconv"
)

func validateBook(book *dto.Book) error {
	if book.Title == "" {
		return errors.New("title is required")
	}

	if book.Author == "" {
		return errors.New("author is required")
	}

	if book.Published <= 0 {
		return errors.New("published year must be greater than 0")
	}

	return nil
}

func getIntId(c echo.Context) (uint, error) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, "Invalid book ID")
	}
	return uint(id), nil
}
