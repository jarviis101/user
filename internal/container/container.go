package container

import (
	"database/sql"
	"user/internal/infrastructure/config"
	"user/internal/infrastructure/database"
)

type Container interface {
	Close() error
}

type container struct {
	connection *sql.DB
}

func NewContainer(config config.Config) (Container, error) {
	connection, err := database.Connect(config.DatabaseDSN)

	if err != nil {
		return nil, err
	}

	return &container{connection}, nil
}

func (c *container) Close() error {
	return c.connection.Close()
}
