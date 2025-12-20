package product

import (
	"errors"
	"fmt"

	"github.com/arya-bhanu/intensive-go/model"
)

func (ps *productService) CreateProduct(product *model.Product) error {
	if product != nil {
		res, err := ps.productStore.CreateOne(*product)
		if err != nil {
			fmt.Printf("error create product: %v\n", err)
			return err
		}
		id, err := res.LastInsertId()
		if err != nil {
			fmt.Printf("unable to get LastInsertId %v\n", err)
			return err
		}
		product.ID = id
		return err
	}
	return errors.New("product is nil")
}

func (ps *productService) GetProduct(id int64) (*model.Product, error) {
	return ps.productStore.FindOne(int(id))
}

func (ps *productService) UpdateProduct(product *model.Product) error {
	if product == nil {
		return errors.New("product is nil")
	}
	return ps.productStore.UpdateProduct(*product)
}

func (ps *productService) DeleteProduct(id int64) error {
	return ps.productStore.DeleteOne("id", id)
}

func (ps *productService) ListProducts(category string) ([]*model.Product, error) {
	return ps.productStore.FindAllInCategories(category)
}

func (ps *productService) BatchUpdateInventory(updates map[int64]int) error {
	return ps.productStore.UpdateTxBulkProductCategoryByID(updates)
}
