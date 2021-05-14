package storage

import "github.com/demogoo/wirego/internal/config"

type CacheConn struct {
	Conf *config.Config
}

func NewCacheConn(conf *config.Config) *CacheConn {
	return &CacheConn{Conf: conf}
}
