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

func TestMemberList(t *testing.T) {
	gin.SetMode(gin.TestMode)

	config := config.NewConfig()
	dbConn := storage.NewDBConn(config)
	service := services.NewMemberService(dbConn)
	handler := NewMemberHandler(service)

	e := gin.Default()
	e.GET("/members", handler.List)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/members", nil)
	e.ServeHTTP(w, r)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	members := make([]*models.Member, 0)
	err := json.Unmarshal(body, &members)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, resp.StatusCode, http.StatusOK)
	assert.Equal(t, members[0], data.Members[0])
}
