package handlers

import (
	"net/http"

	"github.com/demogoo/wirego/internal/services"
	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	service *services.ItemService
}

func NewItemHandler(i *services.ItemService) *ItemHandler {
	return &ItemHandler{service: i}
}

func (i *ItemHandler) List(c *gin.Context) {
	items := i.service.FindAll()
	c.JSON(http.StatusOK, items)
}
