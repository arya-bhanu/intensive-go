package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"github.com/arya-bhanu/intensive-go/db"
	"github.com/arya-bhanu/intensive-go/dep"
	"github.com/arya-bhanu/intensive-go/service"
)

func main() {
	cfg := InitApp()
	if cfg != nil && cfg.DB != nil {
		defer func() {
			cfg.DB.Close()
		}()
	}
	service.NewService(cfg)
}

func InitApp() *dep.Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := db.InitDB()

	if err != nil {
		fmt.Printf("error init db: %v\n", err)
	}

	cfg := dep.Config{
		DB: db,
	}

	return &cfg
}
