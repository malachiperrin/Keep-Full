package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"example.com/keepFull/pkg/db"
	"example.com/keepFull/pkg/products"
)

func main() {

	db.InitDB()
	//products.FetchAllProducts()
	//p, err := products.FetchProductByID(2)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(p)
	//

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//io.WriteString(w, "Test Write")
		t, err := template.ParseGlob("templates/**.html")
		if err != nil {
			log.Fatal(err)
		}

		errs := t.Execute(w, nil)
		if errs != nil {
			log.Fatal(errs)
		}
	})

	http.HandleFunc("/product/{id}", func(w http.ResponseWriter, r *http.Request) {
		// https://pkg.go.dev/net/http#NewServeMux
		var id int
		id = 0
		p, err := products.FetchProductByID(id)

		if err != nil {
			log.Fatal(err)
		}

		io.WriteString(w, p.ProductName)
	})

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		products.FetchAllProducts()
		for _, p := range products.ProductsList {
			io.WriteString(w, p.ProductName)
		}
	})

	http.HandleFunc("/seed", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Seed Route...")
		seeder := db.ProductsSeeder{Table: "products"}
		seeder.Seed()
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
