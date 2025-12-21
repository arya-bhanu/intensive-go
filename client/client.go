package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/arya-bhanu/intensive-go/model"
)

func (c *UserServiceClient) GetUser(ctx context.Context, userID int64) (*model.User, error) {
	resp, err := http.Get(fmt.Sprintf("%s/user/get?id=%d", c.baseURL, userID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	var user model.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *UserServiceClient) ValidateUser(ctx context.Context, userID int64) (bool, error) {
	resp, err := http.Get(fmt.Sprintf("%s/user/validate?id=%d", c.baseURL, userID))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return false, status.Errorf(codes.NotFound, "user not found")
	}

	var result map[string]bool
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	return result["valid"], nil
}

func (c *ProductServiceClient) GetProduct(ctx context.Context, productID int64) (*model.Product, error) {
	// TODO: Implement gRPC client call
	// Hint: make gRPC call to GetProductRPC method
	return nil, status.Errorf(codes.Unimplemented, "client GetProduct not implemented")
}

func (c *ProductServiceClient) CheckInventory(ctx context.Context, productID int64, quantity int32) (bool, error) {
	// TODO: Implement gRPC client call
	// Hint: make gRPC call to CheckInventoryRPC method
	return false, status.Errorf(codes.Unimplemented, "client CheckInventory not implemented")
}
