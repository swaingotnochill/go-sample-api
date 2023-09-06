package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaingotnochill/tempmee/bootstrap"
	"github.com/swaingotnochill/tempmee/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderController struct {
	OrderUseCase domain.OrderUseCase
	Env          *bootstrap.Env
}

func (oc *OrderController) GetAllOrders(c *gin.Context) {
	books, err := oc.OrderUseCase.GetAllOrders(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (oc *OrderController) CreateOrder(c *gin.Context) {
	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	order.ID = primitive.NewObjectID()
	if err := oc.OrderUseCase.CreateOrder(c, &order); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (oc *OrderController) GetOrderByID(c *gin.Context) {
	id := c.Param("id")

	order, err := oc.OrderUseCase.GetOrderByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (oc *OrderController) GetOrdersByCustomerID(c *gin.Context) {
	id := c.Param("id")

	orders, err := oc.OrderUseCase.GetOrdersByCustomerID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}
