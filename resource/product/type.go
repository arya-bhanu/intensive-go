package product

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/arya-bhanu/intensive-go/model"
)

type productStore struct {
	db  *sqlx.DB
	trx *sqlx.Tx
}

type ProductStore interface {
	FindOne(id int) (*model.Product, error)
	CreateOne(product model.Product) (sql.Result, error)
	FindAllInCategories(categories ...string) ([]*model.Product, error)
	UpdateProduct(product model.Product) error
	DeleteOne(colName string, val any) error
	UpdateTxBulkProductCategoryByID(products map[int64]int) error
	UpdateProductOneColumn(filterByColumn string, filterByVal any, columnName string, val any) error
}
