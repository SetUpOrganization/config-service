package clients

import (
	"config_service/internal/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

type UserServiceClient struct {
	Client proto.UsersClient
}

func NewUserServiceClient(address string, timeout time.Duration) (*UserServiceClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("could not connect: %v", err)
	}
	return &UserServiceClient{Client: proto.NewUsersClient(conn)}, nil
}
