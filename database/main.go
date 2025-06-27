package main

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Category struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CategoryID  uuid.UUID `json:"category_id"`
}

func createTable(db *sql.DB) error {
	// Create categories table
	categoryQuery := `
	CREATE TABLE IF NOT EXISTS categories (
		id UUID PRIMARY KEY,
		name TEXT NOT NULL
	);`
	_, err := db.Exec(categoryQuery)
	if err != nil {
		return err
	}

	// Create products table with category reference
	productQuery := `
	CREATE TABLE IF NOT EXISTS products (
		id UUID PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT,
		price NUMERIC(10, 2) NOT NULL,
		category_id UUID,
		FOREIGN KEY (category_id) REFERENCES categories(id)
	);`
	_, err = db.Exec(productQuery)
	return err
}

func newCategory(name string) Category {
	return Category{
		ID:   uuid.New(),
		Name: name,
	}
}

func insertCategory(db *sql.DB, category Category) error {
	stmt, err := db.Prepare("INSERT INTO categories (id, name) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(category.ID, category.Name)
	return err
}

func getCategoryByID(db *sql.DB, id uuid.UUID) (*Category, error) {
	stmt, err := db.Prepare("SELECT id, name FROM categories WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var category Category
	err = stmt.QueryRow(id).Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func getCategories(db *sql.DB) ([]Category, error) {
	rows, err := db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func newProduct(name, description string, price float64, categoryID uuid.UUID) Product {
	return Product{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
	}
}

func insertProduct(db *sql.DB, product Product) error {
	stmt, err := db.Prepare("INSERT INTO products (id, name, description, price, category_id) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Description, product.Price, product.CategoryID)
	return err
}

func updateProduct(db *sql.DB, product Product) error {
	stmt, err := db.Prepare("UPDATE products SET name = $1, description = $2, price = $3, category_id = $4 WHERE id = $5")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Description, product.Price, product.CategoryID, product.ID)
	return err
}

func getProductByID(db *sql.DB, id uuid.UUID) (*Product, error) {
	stmt, err := db.Prepare("SELECT id, name, description, price, category_id FROM products WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var product Product
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func getProductsByCategory(db *sql.DB, categoryID uuid.UUID) ([]Product, error) {
	rows, err := db.Query("SELECT id, name, description, price, category_id FROM products WHERE category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func getProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT id, name, description, price, category_id FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func clearTables(db *sql.DB) error {
	// Delete products first because of foreign key constraint
	_, err := db.Exec("DELETE FROM products")
	if err != nil {
		return err
	}

	// Then delete categories
	_, err = db.Exec("DELETE FROM categories")
	return err
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5433/testdb?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	if err := clearTables(db); err != nil {
		log.Fatalf("Failed to clear tables: %v", err)
	} else {
		log.Println("Tables cleared successfully")
	}

	if err := createTable(db); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	} else {
		log.Println("Tables created successfully")
	}

	// Create categories
	electronics := newCategory("Electronics")
	clothing := newCategory("Clothing")

	if err := insertCategory(db, electronics); err != nil {
		log.Fatalf("Failed to create category: %v", err)
	} else {
		log.Println("Category created successfully:", electronics)
	}

	if err := insertCategory(db, clothing); err != nil {
		log.Fatalf("Failed to create category: %v", err)
	} else {
		log.Println("Category created successfully:", clothing)
	}

	// Create products with categories
	product1 := newProduct("Smartphone", "Latest model smartphone", 699.99, electronics.ID)
	product2 := newProduct("T-Shirt", "Cotton t-shirt", 19.99, clothing.ID)

	if err := insertProduct(db, product1); err != nil {
		log.Fatalf("Failed to create product: %v", err)
	} else {
		log.Println("Product created successfully:", product1)
	}

	if err := insertProduct(db, product2); err != nil {
		log.Fatalf("Failed to create product: %v", err)
	} else {
		log.Println("Product created successfully:", product2)
	}

	// Get products by category
	electronicsProducts, err := getProductsByCategory(db, electronics.ID)
	if err != nil {
		log.Fatalf("Failed to get products by category: %v", err)
	} else {
		log.Printf("Products in %s category: %+v\n", electronics.Name, electronicsProducts)
	}

	// Update product category
	product2.CategoryID = electronics.ID
	if err := updateProduct(db, product2); err != nil {
		log.Fatalf("Failed to update product: %v", err)
	} else {
		log.Println("Product category updated successfully")
	}

	// Retrieve updated product
	updatedProduct, err := getProductByID(db, product2.ID)
	if err != nil {
		log.Fatalf("Failed to retrieve product: %v", err)
	} else {
		// Get the category name for the product
		category, err := getCategoryByID(db, updatedProduct.CategoryID)
		if err != nil {
			log.Fatalf("Failed to get category: %v", err)
		}
		log.Printf("Retrieved product: %+v (Category: %s)\n", updatedProduct, category.Name)
	}
}
