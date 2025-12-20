package product

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/arya-bhanu/intensive-go/model"
)

var (
	ctx = context.Background()
)

func (ps *productStore) UpdateProduct(product model.Product) error {
	query := fmt.Sprintf("UPDATE %s SET name=?, price=?, quantity=?, category=?  WHERE id = ?", ProductTable)
	_, err := ps.db.Exec(query, product.Name, product.Price, product.Quantity, product.Category, product.ID)
	return err
}

func (ps *productStore) UpdateProductOneColumn(filterByColumn string, filterByVal any, columnName string, val any) error {
	query := fmt.Sprintf("UPDATE %s SET %s=? WHERE %s = ?", ProductTable, columnName, filterByColumn)
	result, err := ps.trx.Exec(query, val, filterByVal)
	if err != nil {
		fmt.Printf("error exec update trx: %v\n", err)
		return err
	}
	row, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("error get RowsAffected when update trx: %v\n", err)
		return err
	}
	if row == 0 {
		return errors.New("row not found")
	}
	return err
}

func (ps *productStore) UpdateTxBulkProductCategoryByID(products map[int64]int) error {
	tx, err := ps.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		fmt.Printf("error begin trx: %v\n", err)
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	defer tx.Rollback()
	ps.trx = tx
	for key, val := range products {
		err = ps.UpdateProductOneColumn("id", key, "quantity", val)
		fmt.Printf("error trx update row: %v\n", err)
		if err != nil {
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		fmt.Printf("error commit update bulk: %v\n", err)
		return err
	}
	return err
}
