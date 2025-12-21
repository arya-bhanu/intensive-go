package order

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/arya-bhanu/intensive-go/model"
)

// CreateOrder creates a new order
func (s *OrderService) CreateOrder(ctx context.Context, userID, productID int64, quantity int32) (*model.Order, error) {
	// TODO: Implement this method
	// Hint: 1. Validate user, 2. Get product and check inventory, 3. Create order
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}

// GetOrder retrieves an order by ID
func (s *OrderService) GetOrder(orderID int64) (*model.Order, error) {
	order, exists := s.orders[orderID]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "order not found")
	}
	return order, nil
}
