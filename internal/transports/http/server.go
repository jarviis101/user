package http

import (
	"context"
	"fmt"
	"user/internal/container"
	"user/internal/infrastructure/config"
	"user/internal/transports/http/graphql"
	"user/internal/transports/http/graphql/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/labstack/echo/v4"
)

type Server interface {
	Run() error
	Shutdown(ctx context.Context) error
}

type server struct {
	server    *echo.Echo
	container *container.Container
	config    *config.Config
}

func NewServer(container *container.Container, config *config.Config) Server {
	s := echo.New()

	return &server{s, container, config}
}

func (s *server) Run() error {
	s.injectGraphqlRoutes(s.server)

	return s.server.Start(fmt.Sprintf("%s:%s", s.config.Host, s.config.Port))
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *server) injectGraphqlRoutes(e *echo.Echo) {
	resolver := graph.NewResolver(s.container)
	config := graph.Config{Resolvers: resolver}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(config))
	graphql := graphql.NewGraphqlRouter(e.Group(""), srv, s.config)

	graphql.Populate()
}
