package product

import (
	dep "github.com/arya-bhanu/intensive-go/dep"
	"github.com/arya-bhanu/intensive-go/resource"
)

func NewProductService(cfg *dep.Config) *productService {
	res := resource.NewResource(cfg)
	return &productService{
		productStore: res.ProductStore,
	}
}
