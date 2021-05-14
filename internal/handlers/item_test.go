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

func TestItemList(t *testing.T) {
	gin.SetMode(gin.TestMode)

	conf := config.NewConfig()
	mongoDBConn := storage.NewMongoDBConn(conf)
	service := services.NewItemService(mongoDBConn)
	handler := NewItemHandler(service)

	e := gin.Default()
	e.GET("/items", handler.List)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/items", nil)
	e.ServeHTTP(w, r)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	items := make([]*models.Item, 0)
	err := json.Unmarshal(body, &items)
	if err != nil {
		t.Fatal(err, "=>", string(body))
	}
	assert.Equal(t, resp.StatusCode, http.StatusOK)
	assert.Equal(t, items[0], data.Items[0])
}
