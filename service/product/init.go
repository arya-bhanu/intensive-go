package product

import (
	"github.com/arya-bhanu/intensive-go/model"
)

func NewProductServiceServer() *ProductServiceServer {
	products := map[int64]*model.Product{
		1: {ID: 1, Name: "Laptop", Price: 999.99, Inventory: 10},
		2: {ID: 2, Name: "Phone", Price: 499.99, Inventory: 20},
		3: {ID: 3, Name: "Headphones", Price: 99.99, Inventory: 0},
	}
	return &ProductServiceServer{products: products}
}
