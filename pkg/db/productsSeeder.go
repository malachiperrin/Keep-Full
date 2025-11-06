package db

import (
	"fmt"
)

type ProductsSeeder struct {
	Table string
}

func (ps *ProductsSeeder) Seed() {
	fmt.Println("Seeding Products Table...")
	stmt, err := DB.Query("SELECT * FROM products")
}
