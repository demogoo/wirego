package services

import (
	"github.com/demogoo/wirego/internal/data"
	"github.com/demogoo/wirego/internal/models"
	"github.com/demogoo/wirego/internal/storage"
)

type OrderService struct {
	Cache *storage.CacheConn
}

func NewOrderService(cache *storage.CacheConn) *OrderService {
	return &OrderService{Cache: cache}
}

func (m *OrderService) GetAll() []*models.Order {
	return data.Orders
}
