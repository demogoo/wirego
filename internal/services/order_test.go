package services

import (
	"testing"

	"github.com/demogoo/wirego/internal/config"
	"github.com/demogoo/wirego/internal/data"
	"github.com/demogoo/wirego/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestOrderGet(t *testing.T) {
	conf := config.NewConfig()
	cache := storage.NewCacheConn(conf)
	orderService := NewOrderService(cache)
	orders := orderService.GetAll()
	assert.Equal(t, len(data.Orders), len(orders))
	assert.Equal(t, data.Orders[0], orders[0])
}
