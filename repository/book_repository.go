package repository

import (
	"github.com/emaanmohamed/rest-api/dto"
	"github.com/emaanmohamed/rest-api/models"
)

type BookRepository interface {
	CreateBook(book *dto.Book) error
	GetAllBooks() ([]models.Book, error)
	GetBookByID(id uint) (*models.Book, error)
	UpdateBook(id uint, book *dto.Book) (*models.Book, error)
	DeleteBook(id uint) error
}
