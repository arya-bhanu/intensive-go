package client

import (
	"google.golang.org/grpc"

	"github.com/arya-bhanu/intensive-go/service/product"
	"github.com/arya-bhanu/intensive-go/service/user"
)

func NewUserServiceClient(conn *grpc.ClientConn) user.UserService {
	// Extract address from connection for HTTP calls
	// In a real gRPC implementation, this would use the connection directly
	return &UserServiceClient{baseURL: "http://localhost:50051"}
}

func NewProductServiceClient(conn *grpc.ClientConn) product.ProductService {
	return &ProductServiceClient{conn: conn}
}
