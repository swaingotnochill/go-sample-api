package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaingotnochill/tempmee/bootstrap"
	"github.com/swaingotnochill/tempmee/domain"
)

type HelloController struct {
	HelloUsecase domain.HelloUsecase
	Env          *bootstrap.Env
}

func (hc *HelloController) PrintHello(c *gin.Context) {
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Hello World from server"})
}
