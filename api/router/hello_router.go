package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/swaingotnochill/tempmee/api/controller"
	"github.com/swaingotnochill/tempmee/bootstrap"
	"github.com/swaingotnochill/tempmee/repository"
	"github.com/swaingotnochill/tempmee/usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewHelloRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewHelloRepository(db, "")
	hc := controller.HelloController{
		HelloUsecase: usecase.NewHelloUseCase(ur, timeout),
		Env:          env,
	}
	group.GET("/hello", hc.PrintHello)
}
