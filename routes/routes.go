package routes

import (
	"github.com/emaanmohamed/rest-api/handlers"
	"github.com/emaanmohamed/rest-api/models"
	"github.com/emaanmohamed/rest-api/repository"
	"github.com/labstack/echo/v4"
)

func SetUpRoutes(g *echo.Group) {
	repo := repository.NewGormBookRepository(models.DB)
	bookHandler := handlers.NewBookHandler(repo)
	g.GET("/books", bookHandler.GetAllBooks)
	g.POST("/books", bookHandler.CreateBook)
	g.PUT("/books/:id", bookHandler.UpdateBook)
	g.DELETE("/books/:id", bookHandler.DeleteBook)
}
