package internal

import (
	"user/internal/container"
	"user/internal/infrastructure/config"
	"user/internal/transports/http"
)

type Application interface {
	Run() error
}

type app struct {
	config    *config.Config
	container container.Container
}

func NewApp() (Application, error) {
	config, err := config.LoadConfig()

	if err != nil {
		return nil, err
	}

	container, err := container.NewContainer(*config)

	if err != nil {
		return nil, err
	}

	return &app{config, container}, nil
}

func (a *app) Run() error {
	server := http.NewServer()

	return server.Run()
}
