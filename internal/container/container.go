package container

import (
	"database/sql"
	"user/internal/app/repository"
	"user/internal/app/service"
	"user/internal/infrastructure/config"
	"user/internal/infrastructure/database"
)

type Container struct {
	connection   *sql.DB
	userCreator  service.UserCreator
	userProvider service.UserProvider
}

func NewContainer(config config.Config) (*Container, error) {
	connection, err := database.Connect(config.DatabaseDSN)

	if err != nil {
		return nil, err
	}

	userRepository := repository.NewUserRepository(connection)
	userCreator := service.NewUserCreator(userRepository)
	userProvider := service.NewUserProvider(userRepository)

	return &Container{connection, userCreator, userProvider}, nil
}

func (c *Container) UserCreator() service.UserCreator {
	return c.userCreator
}

func (c *Container) UserProvider() service.UserProvider {
	return c.userProvider
}

func (c *Container) Close() error {
	return c.connection.Close()
}
