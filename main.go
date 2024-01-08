package main

import (
	"database/sql"

	"github.com/marielifm/hexagonal-architecture/app"
	"github.com/marielifm/hexagonal-architecture/app/adapters/db"
)

func main() {
	database, _ := sql.Open("sqlite3", "sqlite.db")
	productDbAdapter := db.NewProductDb(database)
	productService := app.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product 1", 30)
	productService.Enable(product)
}
