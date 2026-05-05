package grpc

import (
	"context"
	"user/internal/app/service"
	"user/proto/user/userpb"
)

type UserService struct {
	userpb.UnimplementedUserServiceServer
	Creator service.UserCreator
}

func (c *UserService) CreateUser(ctx context.Context, payload *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	user, err := c.Creator.Store(ctx, payload.GetFirstName(), payload.GetLastName(), payload.GetEmail(), payload.GetPhone())

	if err != nil {
		return nil, err
	}

	return &userpb.CreateUserResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}, nil
}
