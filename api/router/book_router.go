package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/swaingotnochill/tempmee/api/controller"
	"github.com/swaingotnochill/tempmee/bootstrap"
	"github.com/swaingotnochill/tempmee/domain"
	"github.com/swaingotnochill/tempmee/repository"
	"github.com/swaingotnochill/tempmee/usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewBookRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	br := repository.NewBookRepository(db, domain.CollectionBook)
	bc := controller.BookController{
		BookUseCase: usecase.NewBookUsecase(br, timeout),
		Env:         env,
	}

	group.GET("/books", bc.GetBooks)
	group.GET("/books/:id", bc.GetBookById)
	group.POST("/books", bc.SaveBook)
	group.DELETE("/books/:id", bc.DeleteBook)
}
