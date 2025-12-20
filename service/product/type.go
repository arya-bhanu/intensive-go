package product

import (
	"github.com/arya-bhanu/intensive-go/model"
	productResource "github.com/arya-bhanu/intensive-go/resource/product"
)

type ProductService interface {
	CreateProduct(product *model.Product) error
	GetProduct(id int64) (*model.Product, error)
	UpdateProduct(product *model.Product) error
	DeleteProduct(id int64) error
	ListProducts(category string) ([]*model.Product, error)
	BatchUpdateInventory(updates map[int64]int) error
}

type productService struct {
	productStore productResource.ProductStore
}
