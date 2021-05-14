package storage

import (
	"github.com/demogoo/wirego/internal/config"
)

type DBConn struct {
	Conf *config.Config
}

func NewDBConn(conf *config.Config) *DBConn {
	return &DBConn{Conf: conf}
}
