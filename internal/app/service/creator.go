package service

import (
	"user/internal/app/entity"
	"user/internal/app/repository"
)

type UserCreator interface {
	Store(firstName, lastName, email, phone string) (*entity.User, error)
}

func NewUserCreator(repo repository.UserRepository) UserCreator {
	return &userCreator{repo}
}

type userCreator struct {
	repo repository.UserRepository
}

func (c *userCreator) Store(firstName, lastName, email, phone string) (*entity.User, error) {
	return c.repo.Store(firstName, lastName, email, phone)
}
