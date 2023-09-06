package controller

import (
	"log"
	"net/http"
	"time"

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

// User should be able to create order by just providing book id's in the body. 
// Should we consider providing book id, name and price in the body? 
func (oc *OrderController) CreateOrder(c *gin.Context) {
	var order domain.Order
	var err error
	
	var request domain.OrderRequest
	err = c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	log.Println("Request is ", request.Books)

	if len(request.Books) == 0 {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	order.ID = primitive.NewObjectID()
	customerID := c.GetString("x-user-id")
	order.BookIDs = request.Books
	order.CustomerID, err = primitive.ObjectIDFromHex(customerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	// Calculate the total amount of the order.
	// To do this get the book details from the book ids.
	order.TotalAmount = 0.0
	for _, idString := range order.BookIDs {
		book, err := oc.OrderUseCase.GetBookById(c, idString)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}
		order.TotalAmount += book.Price
	}
	order.CreatedAt = time.Now()
	if err := oc.OrderUseCase.CreateOrder(c, &order); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
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
	id := c.GetString("x-user-id")

	orders, err := oc.OrderUseCase.GetOrdersByCustomerID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}
