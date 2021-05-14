package handlers

import (
	"net/http"

	"github.com/demogoo/wirego/internal/services"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service *services.OrderService
}

func NewOrderHandler(o *services.OrderService) *OrderHandler {
	return &OrderHandler{service: o}
}

func (o *OrderHandler) List(c *gin.Context) {
	orders := o.service.GetAll()
	c.JSON(http.StatusOK, orders)
}

func (o *OrderHandler) PostOrder(c *gin.Context) {
	_, err := c.GetRawData()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
