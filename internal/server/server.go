package server

import (
	"github.com/demogoo/wirego/internal/config"
	"github.com/demogoo/wirego/internal/handlers"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Port   string
	Engine *gin.Engine

	itemHandler   *handlers.ItemHandler
	orderHandler  *handlers.OrderHandler
	memberHandler *handlers.MemberHandler
}

func (s *Server) SetupRoutes() {
	s.Engine.GET("/items", s.itemHandler.List)
	s.Engine.GET("/orders", s.orderHandler.List)
	s.Engine.GET("/members", s.memberHandler.List)
}

func NewServer(
	c *config.Config,
	i *handlers.ItemHandler,
	o *handlers.OrderHandler,
	m *handlers.MemberHandler) *Server {

	return &Server{
		Port:   c.Port,
		Engine: gin.Default(),

		itemHandler:   i,
		orderHandler:  o,
		memberHandler: m,
	}
}
