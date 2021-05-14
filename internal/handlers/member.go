package handlers

import (
	"net/http"

	"github.com/demogoo/wirego/internal/services"
	"github.com/gin-gonic/gin"
)

type MemberHandler struct {
	service *services.MemberService
}

func NewMemberHandler(m *services.MemberService) *MemberHandler {
	return &MemberHandler{service: m}
}

func (i *MemberHandler) List(c *gin.Context) {
	members := i.service.SelectAll()
	c.JSON(http.StatusOK, members)
}
