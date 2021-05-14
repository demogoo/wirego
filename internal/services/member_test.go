package services

import (
	"testing"

	"github.com/demogoo/wirego/internal/config"
	"github.com/demogoo/wirego/internal/data"
	"github.com/demogoo/wirego/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestMemberFindAll(t *testing.T) {
	conf := config.NewConfig()
	db := storage.NewDBConn(conf)
	memberService := NewMemberService(db)
	members := memberService.SelectAll()
	assert.Equal(t, len(data.Members), len(members))
	assert.Equal(t, data.Members[0], members[0])
}
