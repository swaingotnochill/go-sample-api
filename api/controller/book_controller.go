package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaingotnochill/tempmee/bootstrap"
	"github.com/swaingotnochill/tempmee/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookController struct {
	BookUseCase domain.BookUseCase
	Env *bootstrap.Env
}

func (bc *BookController) GetBooks(c *gin.Context) {
	books, err := bc.BookUseCase.GetBooks(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (bc *BookController) GetBookById(c *gin.Context) {
	id := c.Param("id")	

	book, err := bc.BookUseCase.GetBooksByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (bc *BookController) SaveBook(c *gin.Context) {
	var book domain.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	book.ID = primitive.NewObjectID()
	if err := bc.BookUseCase.SaveBook(c, &book); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (bc *BookController) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	if err := bc.BookUseCase.DeleteBook(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
