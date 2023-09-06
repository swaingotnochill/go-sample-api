package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/swaingotnochill/tempmee/api/controller"
	"github.com/swaingotnochill/tempmee/domain"
	"github.com/swaingotnochill/tempmee/domain/mocks"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SetUserID(userID string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("x-user-id", userID)
		c.Next()
	}
}

func TestCreateOrder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		userObjectID := primitive.NewObjectID()
		userID := userObjectID.Hex()

		mockOrderUsecase := new(mocks.OrderUseCase)

		mockOrderUsecase.On("Create", mock.Anything, mock.Anything).Return(nil)

		gin:= gin.Default()
		rec := httptest.NewRecorder()

		oc := controller.OrderController{
			OrderUseCase: mockOrderUsecase,	
		}

		gin.Use(SetUserID(userID))
		gin.POST("/orders", oc.CreateOrder)

		body, err := json.Marshal(domain.SuccessResponse{Message: "Order placed successfully."})
		assert.NoError(t, err)

		bodyString := string(body)

		data := url.Values{}
		data["book_ids"] = []string{"id1", "id2", "id3"}
		
		req := httptest.NewRequest(http.MethodPost, "/orders", strings.NewReader(data.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-url-encoded")
		gin.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, bodyString, rec.Body.String())
		mockOrderUsecase.AssertExpectations(t)
	})
}