package container

import (
	"database/sql"
	"user/internal/app/repository"
	"user/internal/app/service"
	"user/internal/infrastructure/config"
	"user/internal/infrastructure/database"
)

type Container struct {
	connection *sql.DB

	UserCreator service.UserCreator
}

func NewContainer(config config.Config) (*Container, error) {
	connection, err := database.Connect(config.DatabaseDSN)

	if err != nil {
		return nil, err
	}

	userRepository := repository.NewUserRepository(connection)
	userCreator := service.NewUserCreator(userRepository)

	return &Container{connection, userCreator}, nil
}

func (c *Container) Close() error {
	return c.connection.Close()
}
