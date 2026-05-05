package internal

import (
	"fmt"
	"net"
	"user/internal/container"
	"user/internal/infrastructure/config"
	grpc_transport "user/internal/transports/grpc"
	"user/proto/user/userpb"

	grpc_server "google.golang.org/grpc"
)

type grpcApp struct {
	config    *config.Config
	container *container.Container
}

func NewGrpcApp() (Application, error) {
	config, err := config.LoadConfig()

	if err != nil {
		return nil, err
	}

	container, err := container.NewContainer(*config)

	if err != nil {
		return nil, err
	}

	return &grpcApp{config, container}, nil
}

func (a *grpcApp) Run() error {
	server := grpc_server.NewServer()

	userpb.RegisterUserServiceServer(server, &grpc_transport.UserService{
		Creator: a.container.UserCreator(),
	})

	lis, err := net.Listen(
		"tcp",
		fmt.Sprintf("%s:%s", a.config.Host, a.config.Port),
	)

	if err != nil {
		return err
	}

	if err := server.Serve(lis); err != nil {
		return err
	}

	return nil
}
