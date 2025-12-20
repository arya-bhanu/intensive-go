package product

import (
	"github.com/jmoiron/sqlx"
)

func NewProductStore(db *sqlx.DB) ProductStore {
	return &productStore{
		db: db,
	}
}
