package http

import (
	"user/internal/transports/http/graphql"
	"user/internal/transports/http/graphql/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
)

type Server interface {
	Run() error
}

type server struct {
	server *echo.Echo
}

func NewServer() Server {
	s := echo.New()

	return &server{s}
}

func (s *server) Run() error {
	s.injectGraphqlRoutes(s.server)

	return s.server.Start(":8000")
}

func (s *server) injectGraphqlRoutes(e *echo.Echo) {
	config := graph.Config{Resolvers: &graph.Resolver{}}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(config))
	playground := playground.Handler("GraphQL playground", "/query")

	graphql := graphql.NewGraphqlRouter(e.Group(""), srv, playground)

	graphql.Populate()
}
