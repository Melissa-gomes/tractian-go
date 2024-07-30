package controller

import (
	"tractian-go/internal/service"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service service.Service
}

func NewController(serv service.Service) *Controller {
	return &Controller{
		Service: serv,
	}
}

func (c Controller) List(ctx *gin.Context) {
	_, err := c.Service.List(ctx)
	if err != nil {

	}

}
