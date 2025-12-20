package product

import (
	"database/sql"
	"fmt"

	"github.com/arya-bhanu/intensive-go/model"
)

func (ps *productStore) CreateOne(product model.Product) (sql.Result, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, price, quantity, category) VALUES (?, ?, ?, ?)", ProductTable)
	result, err := ps.db.Exec(query, product.Name, product.Price, product.Quantity, product.Category)
	return result, err
}
