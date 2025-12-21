package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/arya-bhanu/intensive-go/service/order"
)

func ConnectToServices(userServiceAddr, productServiceAddr string) (*order.OrderService, error) {
	// TODO: Implement this function
	// Hint: create gRPC connections with interceptors, create clients, return OrderService
	return nil, status.Errorf(codes.Unimplemented, "ConnectToServices not implemented")
}
