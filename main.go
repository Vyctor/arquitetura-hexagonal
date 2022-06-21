package main

import (
	"database/sql"

	db2 "github.com/codeedu/go-hexagonal/adapters/db"
	"github.com/codeedu/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDB(db)
	productService := application.NewProductService(productDbAdapter)

	product, err := productService.Create("Product Test", 30.0)

	if err != nil {
		panic(err)
	}

	productService.Enable(product)
}
