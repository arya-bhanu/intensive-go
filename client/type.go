package client

import "google.golang.org/grpc"

type UserServiceClient struct {
	baseURL string
}

type ProductServiceClient struct {
	conn *grpc.ClientConn
}
