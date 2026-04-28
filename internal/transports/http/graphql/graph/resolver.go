package graph

import "user/internal/app/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.
//

//go:generate go run github.com/99designs/gqlgen generate

type ResolverService interface {
	UserCreator() service.UserCreator
}

type Resolver struct {
	services ResolverService
}

func NewResolver(services ResolverService) *Resolver {
	return &Resolver{services}
}
