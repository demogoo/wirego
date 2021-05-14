package services

import (
	"testing"

	"github.com/demogoo/wirego/internal/config"
	"github.com/demogoo/wirego/internal/data"
	"github.com/demogoo/wirego/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestItemFindAll(t *testing.T) {
	conf := config.NewConfig()
	mongoDB := storage.NewMongoDBConn(conf)
	itemService := NewItemService(mongoDB)
	items := itemService.FindAll()
	assert.Equal(t, len(data.Items), len(items))
	assert.Equal(t, data.Items[0], items[0])
}
