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

func NewOrderRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	or := repository.NewOrderRepository(db, domain.CollectionOrders)
	oc := controller.OrderController{
		OrderUseCase: usecase.NewOrderUseCase(or, timeout),
		Env:          env,
	}

	group.GET("/orders/:id", oc.GetOrderByID)
	group.GET("/orders/history", oc.GetOrdersByCustomerID)
	group.POST("/orders", oc.CreateOrder)
	group.GET("/orders", oc.GetAllOrders)
}
