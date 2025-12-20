package resource

import (
	dep "github.com/arya-bhanu/intensive-go/dep"
	"github.com/arya-bhanu/intensive-go/resource/product"
)

func NewResource(cfg *dep.Config) *Resource {
	productStore := product.NewProductStore(cfg.DB)
	return &Resource{
		ProductStore: productStore,
	}
}
