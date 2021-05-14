package services

import (
	"github.com/demogoo/wirego/internal/data"
	"github.com/demogoo/wirego/internal/models"
	"github.com/demogoo/wirego/internal/storage"
)

type MemberService struct {
	DB *storage.DBConn
}

func NewMemberService(db *storage.DBConn) *MemberService {
	return &MemberService{DB: db}
}

func (m *MemberService) SelectAll() []*models.Member {
	return data.Members
}
