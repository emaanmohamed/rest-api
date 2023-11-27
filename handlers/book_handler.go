package handlers

import (
	"github.com/emaanmohamed/rest-api/dto"
	"github.com/emaanmohamed/rest-api/models"
	"github.com/emaanmohamed/rest-api/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type BookHandler struct {
	Repository repository.BookRepository
}

func NewBookHandler(repo repository.BookRepository) *BookHandler {
	return &BookHandler{Repository: repo}
}

func (h *BookHandler) CreateBook(c echo.Context) error {
	book := new(dto.Book)
	if err := c.Bind(book); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	// validate book
	if err := validateBook(book); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// create book
	err := h.Repository.CreateBook(book)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create book")
	}
	return c.JSON(http.StatusCreated, book)
}

func (h *BookHandler) GetAllBooks(c echo.Context) error {
	books, err := h.Repository.GetAllBooks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get books")
	}
	return c.JSON(http.StatusOK, books)
}

func (h *BookHandler) UpdateBook(c echo.Context) error {
	id, err := getIntId(c)
	if err != nil {
		log.Error("Invalid book ID: ", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid book ID")
	}
	// Create a new dto.Book instance and bind the request data to it
	var updatedBook dto.Book
	if err := c.Bind(&updatedBook); err != nil {
		log.Error("Invalid request data ", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid request data")
	}

	book := new(models.Book)
	if book, err = h.Repository.UpdateBook(id, &updatedBook); err != nil {
		log.Error("Failed to update the book ", err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to update the book.")
	}

	return c.JSON(http.StatusOK, book)
}

func (h *BookHandler) DeleteBook(c echo.Context) error {
	id, err := getIntId(c)
	if err != nil {
		log.Error("Invalid book ID: ", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid book ID")
	}

	if err := h.Repository.DeleteBook(id); err != nil {
		log.Error("Book not found ", err.Error())
		return c.JSON(http.StatusNotFound, "Book not found")
	}

	return c.NoContent(http.StatusNoContent)
}
