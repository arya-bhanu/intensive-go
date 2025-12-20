package product

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/arya-bhanu/intensive-go/model"
	"github.com/arya-bhanu/intensive-go/utils/resource"
)

func (ps *productStore) FindOne(id int) (*model.Product, error) {
	query := fmt.Sprintf("SELECT id, name, price, quantity, category FROM %s WHERE id = ?", ProductTable)
	product := model.Product{}
	err := ps.db.Get(&product, query, id)
	return &product, err
}

func (ps *productStore) FindAllInCategories(categories ...string) ([]*model.Product, error) {
	categoriesFormattedQuery := resource.CreateQueryQuestionMarks(categories)
	query := fmt.Sprintf("SELECT id, name, price, quantity, category FROM %s WHERE category IN (%s)", ProductTable, categoriesFormattedQuery)
	query, args, err := sqlx.In(query, categories)
	if err != nil {
		return nil, err
	}
	products := []*model.Product{}
	err = ps.db.Select(&products, query, args...)
	return products, err
}
