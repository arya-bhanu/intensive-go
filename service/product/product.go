package product

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	model "github.com/arya-bhanu/intensive-go/model"
)

// GetProduct retrieves a product by ID
func (s *ProductServiceServer) GetProduct(ctx context.Context, productID int64) (*model.Product, error) {
	// TODO: Implement this method
	// Hint: check if product exists, return product or gRPC NotFound error
	prod, ok := s.products[productID]
	if !ok {
		return nil, status.Error(codes.NotFound, "product not found")
	}
	if prod == nil {
		return nil, status.Error(codes.Unavailable, "product is nil")
	}
	return prod, nil
}

// CheckInventory checks if a product is available in the requested quantity
func (s *ProductServiceServer) CheckInventory(ctx context.Context, productID int64, quantity int32) (bool, error) {
	// TODO: Implement this method
	// Hint: check product exists and has sufficient inventory
	prod, err := s.GetProduct(ctx, productID)
	if err != nil {
		return false, err
	}
	return prod.Inventory <= quantity, nil
}

func (s *ProductServiceServer) GetProductRPC(ctx context.Context, req *GetProductRequest) (*GetProductResponse, error) {
	product, err := s.GetProduct(ctx, req.ProductId)
	if err != nil {
		return nil, err
	}
	return &GetProductResponse{Product: product}, nil
}

func (s *ProductServiceServer) CheckInventoryRPC(ctx context.Context, req *CheckInventoryRequest) (*CheckInventoryResponse, error) {
	available, err := s.CheckInventory(ctx, req.ProductId, req.Quantity)
	if err != nil {
		return nil, err
	}
	return &CheckInventoryResponse{Available: available}, nil
}
