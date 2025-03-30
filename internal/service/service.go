package service

import (
	"config_service/internal/clients"
	"config_service/internal/repo"
	"time"
)

type Service struct {
	repo *repo.Repo
	uCli *clients.UserServiceClient
}

func NewService(repo *repo.Repo, userAddress string, timeout time.Duration) (*Service, error) {
	userClient, err := clients.NewUserServiceClient(userAddress, timeout)
	if err != nil {
		return nil, err
	}

	return &Service{
		repo: repo,
		uCli: userClient,
	}, nil
}
