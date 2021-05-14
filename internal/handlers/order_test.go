package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/demogoo/wirego/internal/config"
	"github.com/demogoo/wirego/internal/data"
	"github.com/demogoo/wirego/internal/models"
	"github.com/demogoo/wirego/internal/services"
	"github.com/demogoo/wirego/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestOrderList(t *testing.T) {
	gin.SetMode(gin.TestMode)

	config := config.NewConfig()
	cache := storage.NewCacheConn(config)
	service := services.NewOrderService(cache)
	handler := NewOrderHandler(service)

	e := gin.Default()
	e.GET("/orders", handler.List)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/orders", nil)
	e.ServeHTTP(w, r)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	orders := make([]*models.Order, 0)
	err := json.Unmarshal(body, &orders)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, resp.StatusCode, http.StatusOK)
	assert.Equal(t, orders[0], data.Orders[0])
}
