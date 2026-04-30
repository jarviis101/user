package graphql

import (
	"user/internal/infrastructure/config"
	"user/internal/transports"
	"user/internal/transports/http/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
)

type router struct {
	group  *echo.Group
	server *handler.Server
	config *config.Config
}

func NewGraphqlRouter(g *echo.Group, srv *handler.Server, config *config.Config) transports.HttpRouter {
	return &router{g, srv, config}
}

func (r *router) Populate() {
	r.group.Add("POST", "/query", r.query, middleware.Auth(r.config.Secret))

	if r.config.Debug {
		r.group.Add("GET", "/graphql", r.playground)
	}
}

func (r *router) query(c echo.Context) error {
	r.server.ServeHTTP(c.Response(), c.Request())

	return nil
}

func (r *router) playground(c echo.Context) error {
	playground := playground.Handler("GraphQL playground", "/query")
	playground.ServeHTTP(c.Response(), c.Request())

	return nil
}
