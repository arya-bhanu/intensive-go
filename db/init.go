package db

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func InitDB() (*sqlx.DB, error) {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	db_name := os.Getenv("MYSQL_DB_NAME")
	host := os.Getenv("MYSQL_HOST")

	db_url := fmt.Sprintf("%s:%s@%s/%s", user, password, host, db_name)

	db, err := sqlx.Open("mysql", db_url)

	if err != nil {
		fmt.Printf("error open connection db %v\n", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("error ping db: %v\n", err)
		return nil, err
	}
	return db, err
}
