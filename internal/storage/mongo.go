package storage

import "github.com/demogoo/wirego/internal/config"

type MongoDBConn struct {
	Conf *config.Config
}

func NewMongoDBConn(conf *config.Config) *MongoDBConn {
	return &MongoDBConn{Conf: conf}
}
