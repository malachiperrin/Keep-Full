package products

import (
	"database/sql"
	"fmt"
	"log"

	"example.com/keepFull/pkg/db"
)

type Product struct {
	ProductID           int
	ProductName         string
	ProductDependencies []sql.Null[string]
	VendorID            int
}

func FetchAllProducts() {
	rows, err := db.DB.Query("SELECT * FROM Products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var dependencies sql.Null[string]
		var vendor int
		err := rows.Scan(&id, &name, &dependencies, &vendor)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

}

func FetchProductByID(id int) (*Product, error) {
	stmt, err := db.DB.Prepare("SELECT * FROM Products WHERE product_id = ?")

	if err != nil {
		fmt.Println("Error on line 40 - products.go: failed to prepare query...")
		log.Fatal(err.Error())
	}

	defer stmt.Close()
	p := new(Product)

	// Idiomatic Go: inline variable declaration and error check
	//
	// This line:
	//     if r, err := stmt.Query(id); err != nil
	//
	// Does the following in a concise form:
	// 1. Declares and assigns two variables:
	//        r   - the result of stmt.Query(id) (a *Rows value)
	//        err - the error returned by stmt.Query(id)
	// 2. Immediately checks if 'err' is not nil.
	//    - If err != nil, the 'if' block runs (usually to handle the error).
	//    - If err == nil, 'r' is available inside the 'if' block or any attached 'else'.
	//
	// It's equivalent to writing:
	//
	//     r, err := stmt.Query(id)
	//     if err != nil {
	//         // handle error
	//     } else {
	//         // use r
	//     }
	//
	// But the idiomatic form is cleaner and scopes 'r' and 'err' to just the 'if' block.

	// Idiomatic Go: handle the error early, return, and keep success path unindented.
	// Avoid wrapping the success path in an 'else' block.
	// This makes the code flatter, more readable, and easier to maintain.

	r, err := stmt.Query(id)
	if err != nil {
		return p, err
	}
	defer r.Close()

	// process r here
	r.Next()
	var val int
	var val2 string
	var val3 sql.Null[string]
	var val4 int
	r.Scan(&val, &val2, &val3)
	p.ProductID = val
	p.ProductName = val2
	p.ProductDependencies = append(p.ProductDependencies, val3)
	p.VendorID = val4
	return p, nil

}
