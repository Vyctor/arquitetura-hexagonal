package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/codeedu/go-hexagonal/adapters/db"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products(
			id string, 
			name string, 
			price float, 
			status string
		);`

	stmt, err := db.Prepare(table)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products(id, name, price, status) VALUES("abc", "Product Test", 0, "disabled");`

	stmt, err := db.Prepare(insert)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDb_Get(test *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDB(Db)
	product, error := productDb.Get("abc")

	require.Nil(test, error)
	require.Equal(test, "Product Test", product.GetName())
	require.Equal(test, 0.0, product.GetPrice())
	require.Equal(test, "disabled", product.GetStatus())
}
