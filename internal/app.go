package internal

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"user/internal/container"
	"user/internal/infrastructure/config"
	httpTransport "user/internal/transports/http"
)

type Application interface {
	Run() error
}

type app struct {
	config    *config.Config
	container *container.Container
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
	server := httpTransport.NewServer(a.container)

	serverError := make(chan error, 1)
	go func() {
		if err := server.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverError <- err
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverError:
		return err
	case <-quit:
	}

	log.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Http server shutdown error: %v", err)
	}

	if err := a.container.Close(); err != nil {
		log.Printf("Database close error: %v", err)
	}

	return nil
}
