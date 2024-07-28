package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host         = "localhost"
	port         = 5432
	databaseName = "mydatabase"
	username     = "myuser"
	password     = "mypassword"
)

var db *sql.DB

type Product struct {
	ID    int
	Name  string
	Price int
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, databaseName)

	log.Printf("Connecting to database with connection string: %s\n", psqlInfo)

	sdb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database connection: %v\n", err)
	}

	db = sdb

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v\n", err)
	}

	log.Println("Successfully connected to the database")

	// Example usage
	// err = createProduct(&Product{Name: "GO FIRST", Price: 123})
	// if err != nil {
	// 	log.Fatalf("Error creating product: %v\n", err)
	// }

	// product, err := getProduct(1)
	// if err != nil {
	// 	log.Fatalf("Error getting product: %v\n", err)
	// }
	// log.Printf("Product: %+v\n", product)

	// err = updateProduct(1, &Product{Name: "GO BOOK", Price: 444})
	// if err != nil {
	// 	log.Fatalf("Error updating product: %v\n", err)
	// }

	// err = deleteProduct(2)
	// if err != nil {
	// 	log.Fatalf("Error deleting product: %v\n", err)
	// }

	products, err := getProducts()
	if err != nil {
		log.Fatalf("Error getting products: %v\n", err)
	}
	log.Printf("Products: %+v\n", products)
}

func createProduct(product *Product) error {
	if productExistsByName(product.Name) {
		return errors.New("product name already exists")
	}
	_, err := db.Exec(
		"INSERT INTO public.products(name, price) VALUES ($1, $2);",
		product.Name,
		product.Price,
	)
	return err
}

func getProduct(id int) (Product, error) {
	var p Product
	row := db.QueryRow("SELECT id, name, price FROM products WHERE id=$1;", id)
	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err == sql.ErrNoRows {
		return Product{}, errors.New("product not found")
	} else if err != nil {
		return Product{}, err
	}
	return p, nil
}

func updateProduct(id int, product *Product) error {
	if !productExistsByID(id) {
		return errors.New("product not found")
	}
	if productExistsByNameExceptID(product.Name, id) {
		return errors.New("product name already exists")
	}
	_, err := db.Exec(
		"UPDATE public.products SET name=$2, price=$3 WHERE id=$1;", id,
		product.Name,
		product.Price,
	)
	return err
}

func deleteProduct(id int) error {
	if !productExistsByID(id) {
		return errors.New("product not found")
	}
	_, err := db.Exec("DELETE FROM public.products WHERE id=$1;", id)
	return err
}

func getProducts() ([]Product, error) {
	//pointer rows from Query
	rows, err := db.Query("SELECT id, name, price FROM products;")
	if err != nil {
		return nil, err
	}
	//defer for last process use for close rows
	defer rows.Close()

	//variable slice products for collect products data from database
	var products []Product
	//loop rows เพื่อยัดลง slice products
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func productExistsByID(id int) bool {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM products WHERE id=$1);", id).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("Error checking if product exists by ID: %v\n", err)
	}
	return exists
}

func productExistsByName(name string) bool {
	var exists bool
	//select exists is subQuery return boolean if it found
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM products WHERE name=$1);", name).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("Error checking if product exists by name: %v\n", err)
	}
	return exists
}

func productExistsByNameExceptID(name string, id int) bool {
	var exists bool
	//select exists is subQuery return boolean if it found
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM products WHERE name=$1 AND id!=$2);", name, id).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("Error checking if product exists by name except ID: %v\n", err)
	}
	return exists
}
