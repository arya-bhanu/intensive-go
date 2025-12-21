package order

import (
	"github.com/arya-bhanu/intensive-go/model"
	"github.com/arya-bhanu/intensive-go/service/product"
	"github.com/arya-bhanu/intensive-go/service/user"
)

func NewOrderService(userClient user.UserService, productClient product.ProductService) *OrderService {
	return &OrderService{
		userClient:    userClient,
		productClient: productClient,
		orders:        make(map[int64]*model.Order),
		nextOrderID:   1,
	}
}
