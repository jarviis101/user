package service

import (
	"context"
	"user/internal/app/entity"
	"user/internal/app/repository"
)

type UserCreator interface {
	Store(ctx context.Context, firstName, lastName, email, phone string) (*entity.User, error)
}

func NewUserCreator(repo repository.UserRepository) UserCreator {
	return &userCreator{repo}
}

type userCreator struct {
	repo repository.UserRepository
}

func (c *userCreator) Store(ctx context.Context, firstName, lastName, email, phone string) (*entity.User, error) {
	return c.repo.Store(ctx, firstName, lastName, email, phone)
}
