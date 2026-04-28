package http

import (
	"context"
	"user/internal/container"
	"user/internal/transports/http/graphql"
	"user/internal/transports/http/graphql/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
)

type Server interface {
	Run() error
	Shutdown(ctx context.Context) error
}

type server struct {
	server    *echo.Echo
	container *container.Container
}

func NewServer(container *container.Container) Server {
	s := echo.New()

	return &server{s, container}
}

func (s *server) Run() error {
	s.injectGraphqlRoutes(s.server)

	return s.server.Start(":8000")
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *server) injectGraphqlRoutes(e *echo.Echo) {
	resolver := graph.NewResolver(s.container)
	config := graph.Config{Resolvers: resolver}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(config))
	playground := playground.Handler("GraphQL playground", "/query")

	graphql := graphql.NewGraphqlRouter(e.Group(""), srv, playground)

	graphql.Populate()
}
