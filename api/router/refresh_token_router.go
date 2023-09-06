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

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}