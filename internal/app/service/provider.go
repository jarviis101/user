package service

import (
	"context"
	"user/internal/app/entity"
	"user/internal/app/repository"
)

type UserProvider interface {
	ByCriteria(ctx context.Context, filter entity.UserFilter) ([]entity.User, error)
}

type userProvider struct {
	repo repository.UserRepository
}

func NewUserProvider(repo repository.UserRepository) UserProvider {
	return &userProvider{repo: repo}
}

func (f *userProvider) ByCriteria(ctx context.Context, filter entity.UserFilter) ([]entity.User, error) {
	return f.repo.FindByCriteria(ctx, filter)
}
