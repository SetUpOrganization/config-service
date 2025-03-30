package clients

import (
	"context"
	"fmt"
	"time"
)

type UserServiceClient struct {
	Client proto.UserServiceClient // TODO: закинуть proto и тд
}

func NewUserServiceClient(address string, timeout time.Duration) (*UserServiceClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("could not connect: %v", err)
	}
	return &UserServiceClient{Client: proto.NewUserServiceClient(conn)}, nil
}
