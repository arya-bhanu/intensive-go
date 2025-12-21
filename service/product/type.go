package product

import (
	"context"

	"github.com/arya-bhanu/intensive-go/model"
)

type ProductService interface {
	GetProduct(ctx context.Context, productID int64) (*model.Product, error)
	CheckInventory(ctx context.Context, productID int64, quantity int32) (bool, error)
}

type ProductServiceServer struct {
	products map[int64]*model.Product
}

type GetProductRequest struct {
	ProductId int64 `json:"product_id"`
}

type GetProductResponse struct {
	Product *model.Product `json:"product"`
}

type CheckInventoryRequest struct {
	ProductId int64 `json:"product_id"`
	Quantity  int32 `json:"quantity"`
}

type CheckInventoryResponse struct {
	Available bool `json:"available"`
}
