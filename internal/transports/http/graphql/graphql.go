package graphql

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/labstack/echo/v4"
)

type Router interface {
	Populate()
}

type router struct {
	group  *echo.Group
	server *handler.Server
	pg     http.HandlerFunc
}

func NewGraphqlRouter(g *echo.Group, srv *handler.Server, pg http.HandlerFunc) Router {
	return &router{g, srv, pg}
}

func (r *router) Populate() {
	r.group.Add("GET", "/graphql", r.playground)
	r.group.Add("POST", "/query", r.query)
}

func (r *router) query(c echo.Context) error {
	r.server.ServeHTTP(c.Response(), c.Request())

	return nil
}

func (r *router) playground(c echo.Context) error {
	r.pg.ServeHTTP(c.Response(), c.Request())

	return nil
}
