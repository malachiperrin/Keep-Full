package main

import (
	"example.com/keepFull/pkg/db"
	"example.com/keepFull/pkg/products"
)

func main() {
	db.InitDB()
	products.FetchAllProducts()
	//p, err := products.FetchProductByID(2)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(p)
}
