package service

import (
	"github.com/arya-bhanu/intensive-go/dep"
	"github.com/arya-bhanu/intensive-go/service/product"
)

func NewService(cfg *dep.Config) *Service {
	productService := product.NewProductService(cfg)
	return &Service{
		ProductService: productService,
	}
}
