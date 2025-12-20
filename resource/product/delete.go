package product

import "fmt"

func (ps *productStore) DeleteOne(colName string, val any) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", ProductTable, colName)
	_, err := ps.db.Exec(query, val)
	return err
}
