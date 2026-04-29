package service

import (
	"user/internal/app/entity"
	"user/internal/app/repository"
)

type UserProvider interface {
	ByCriteria(filter entity.UserFilter) ([]entity.User, error)
}

type userProvider struct {
	repo repository.UserRepository
}

func NewUserProvider(repo repository.UserRepository) UserProvider {
	return &userProvider{repo: repo}
}

func (f *userProvider) ByCriteria(filter entity.UserFilter) ([]entity.User, error) {
	return f.repo.FindByCriteria(filter)
}
