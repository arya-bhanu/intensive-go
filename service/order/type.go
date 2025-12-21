package order

import (
	"github.com/arya-bhanu/intensive-go/model"
	"github.com/arya-bhanu/intensive-go/service/product"
	"github.com/arya-bhanu/intensive-go/service/user"
)

type OrderService struct {
	userClient    user.UserService
	productClient product.ProductService
	orders        map[int64]*model.Order
	nextOrderID   int64
}
