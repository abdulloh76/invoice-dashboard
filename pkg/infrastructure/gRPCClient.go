package infrastructure

import (
	"context"

	"github.com/abdulloh76/invoice-dashboard/pkg/infrastructure/userGrpc"
	"github.com/abdulloh76/invoice-dashboard/pkg/types"

	"google.golang.org/grpc"
)

type UserGrpcClient struct {
	client userGrpc.UserClient
}

func NewUserGrpcClient(userGrpcPort string) *UserGrpcClient {
	conn, err := grpc.Dial(":"+userGrpcPort, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	c := userGrpc.NewUserClient(conn)

	return &UserGrpcClient{
		client: c,
	}
}

func (c *UserGrpcClient) GetUserAddress(ctx context.Context, userId string) (*types.AddressModel, error) {
	address, err := c.client.GetUserAddress(ctx, &userGrpc.GetRequest{Id: userId})

	if err != nil {
		panic(err)
	}

	senderAddress := &types.AddressModel{
		Street:   address.Street,
		City:     address.City,
		PostCode: address.PostCode,
		Country:  address.Country,
	}

	return senderAddress, err
}
