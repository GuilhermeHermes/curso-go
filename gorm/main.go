package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// User has many CreditCards, UserID is the foreign key
type User struct {
	gorm.Model
	Name       string
	Email      string `gorm:"uniqueIndex"`
	Age        uint8
	Profile    Profile      // User has one Profile
	CreditCard []CreditCard // User has many CreditCards
	Languages  []Language   `gorm:"many2many:user_languages;"`
	Orders     []Order      // User has many Orders
}

// Profile belongs to User, UserID is the foreign key
type Profile struct {
	gorm.Model
	UserID      uint
	Bio         string
	PhoneNumber string
	Address     string
}

// CreditCard belongs to User, UserID is the foreign key
type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

// Language belongs to many Users
type Language struct {
	gorm.Model
	Name  string
	Users []User `gorm:"many2many:user_languages;"`
}

// Order belongs to User
type Order struct {
	gorm.Model
	UserID      uint
	OrderNumber string
	Total       float64
	Status      string
	Items       []OrderItem // Order has many OrderItems
}

// OrderItem belongs to Order
type OrderItem struct {
	gorm.Model
	OrderID   uint
	ProductID uint
	Quantity  int
	Price     float64
	Product   Product `gorm:"foreignKey:ProductID"`
}

// Product belongs to Category
type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	CategoryID  uint
	Category    Category
}

// Category has many Products
type Category struct {
	gorm.Model
	Name        string
	Description string
	Products    []Product
}

func main() {
	// Connect to database
	dsn := "host=localhost user=postgres password=postgres dbname=testdb port=5433 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate schemas
	err = db.AutoMigrate(
		&User{},
		&Profile{},
		&CreditCard{},
		&Language{},
		&Category{},  // Migrate Category first
		&Product{},   // Then Product which depends on Category
		&Order{},     // Then Order
		&OrderItem{}, // Finally OrderItem which depends on Product and Order
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Demonstration of different operations and relationships
	demonstrateCRUD(db)
	demonstrateHasOne(db)
	demonstrateHasMany(db)
	demonstrateBelongsTo(db)
	demonstrateManyToMany(db)
	demonstrateComplexRelationships(db)
}

func demonstrateCRUD(db *gorm.DB) {
	fmt.Println("\n=== CRUD Operations ===")

	// Create
	user := User{
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   30,
	}
	result := db.Create(&user)
	if result.Error != nil {
		log.Printf("Error creating user: %v", result.Error)
	} else {
		fmt.Printf("Created user: %v, ID: %v\n", user.Name, user.ID)
	}

	// Read
	var retrievedUser User
	result = db.First(&retrievedUser, user.ID)
	if result.Error != nil {
		log.Printf("Error retrieving user: %v", result.Error)
	} else {
		fmt.Printf("Retrieved user: %v, Email: %v\n", retrievedUser.Name, retrievedUser.Email)
	}

	// Update
	db.Model(&retrievedUser).Updates(User{Name: "John Updated", Age: 31})
	fmt.Printf("Updated user: %v, Age: %v\n", retrievedUser.Name, retrievedUser.Age)

	// Delete
	// Using unscoped to permanently delete, otherwise GORM uses soft delete
	// db.Unscoped().Delete(&retrievedUser)
	// fmt.Println("Deleted user")

	// For demonstration purposes, let's not delete the user
	fmt.Println("User not deleted for demonstration purposes")
}

func demonstrateHasOne(db *gorm.DB) {
	fmt.Println("\n=== Has One Relationship ===")

	// Create a user
	user := User{
		Name:  "Alice",
		Email: "alice@example.com",
		Age:   25,
	}
	db.Create(&user)

	// Create a profile for the user (HasOne relationship)
	profile := Profile{
		UserID:      user.ID,
		Bio:         "Software developer",
		PhoneNumber: "123-456-7890",
		Address:     "123 Main St",
	}
	db.Create(&profile)

	// Retrieve user with profile
	var userWithProfile User
	db.Preload("Profile").First(&userWithProfile, user.ID)
	fmt.Printf("User: %v, Profile Bio: %v\n", userWithProfile.Name, userWithProfile.Profile.Bio)
}

func demonstrateHasMany(db *gorm.DB) {
	fmt.Println("\n=== Has Many Relationship ===")

	// Create a user
	user := User{
		Name:  "Bob",
		Email: "bob@example.com",
		Age:   35,
	}
	db.Create(&user)

	// Create credit cards for the user (HasMany relationship)
	creditCards := []CreditCard{
		{Number: "1111-2222-3333-4444", UserID: user.ID},
		{Number: "5555-6666-7777-8888", UserID: user.ID},
	}
	for _, card := range creditCards {
		db.Create(&card)
	}

	// Retrieve user with credit cards
	var userWithCards User
	db.Preload("CreditCard").First(&userWithCards, user.ID)
	fmt.Printf("User: %v has %v credit cards\n", userWithCards.Name, len(userWithCards.CreditCard))
	for i, card := range userWithCards.CreditCard {
		fmt.Printf("  Card %d: %v\n", i+1, card.Number)
	}
}

func demonstrateBelongsTo(db *gorm.DB) {
	fmt.Println("\n=== Belongs To Relationship ===")

	// Create a category
	category := Category{
		Name:        "Electronics",
		Description: "Electronic devices and gadgets",
	}
	db.Create(&category)

	// Create a product that belongs to the category
	product := Product{
		Name:        "Smartphone",
		Description: "Latest model",
		Price:       999.99,
		CategoryID:  category.ID,
	}
	db.Create(&product)

	// Retrieve product with its category
	var retrievedProduct Product
	db.Preload("Category").First(&retrievedProduct, product.ID)
	fmt.Printf("Product: %v belongs to Category: %v\n", retrievedProduct.Name, retrievedProduct.Category.Name)
}

func demonstrateManyToMany(db *gorm.DB) {
	fmt.Println("\n=== Many to Many Relationship ===")

	// Create a user
	user := User{
		Name:  "Charlie",
		Email: "charlie@example.com",
		Age:   28,
	}
	db.Create(&user)

	// Create languages
	languages := []Language{
		{Name: "Go"},
		{Name: "Python"},
		{Name: "JavaScript"},
	}
	for i := range languages {
		db.Create(&languages[i])
	}

	// Associate languages with user
	db.Model(&user).Association("Languages").Append(&languages)

	// Retrieve user with languages
	var userWithLanguages User
	db.Preload("Languages").First(&userWithLanguages, user.ID)
	fmt.Printf("User: %v knows %v languages\n", userWithLanguages.Name, len(userWithLanguages.Languages))
	for i, lang := range userWithLanguages.Languages {
		fmt.Printf("  Language %d: %v\n", i+1, lang.Name)
	}

	// Retrieve languages with users
	var goLang Language
	db.Preload("Users").Where("name = ?", "Go").First(&goLang)
	fmt.Printf("Language: %v is known by %v users\n", goLang.Name, len(goLang.Users))
}

func demonstrateComplexRelationships(db *gorm.DB) {
	fmt.Println("\n=== Complex Relationships ===")

	// Create a user
	user := User{
		Name:  "David",
		Email: "david@example.com",
		Age:   40,
	}
	db.Create(&user)

	// Create categories
	categories := []Category{
		{Name: "Books", Description: "Physical and digital books"},
		{Name: "Clothing", Description: "Apparel and accessories"},
	}
	for i := range categories {
		db.Create(&categories[i])
	}

	// Create products
	products := []Product{
		{Name: "Go Programming", Description: "Learn Go programming", Price: 49.99, CategoryID: categories[0].ID},
		{Name: "T-Shirt", Description: "Cotton t-shirt", Price: 19.99, CategoryID: categories[1].ID},
		{Name: "Hoodie", Description: "Warm hoodie", Price: 39.99, CategoryID: categories[1].ID},
	}
	for i := range products {
		db.Create(&products[i])
	}

	// Create an order with items
	order := Order{
		UserID:      user.ID,
		OrderNumber: fmt.Sprintf("ORD-%v", time.Now().Unix()),
		Total:       109.97,
		Status:      "pending",
	}
	db.Create(&order)

	// Create order items
	orderItems := []OrderItem{
		{OrderID: order.ID, ProductID: products[0].ID, Quantity: 1, Price: products[0].Price},
		{OrderID: order.ID, ProductID: products[1].ID, Quantity: 2, Price: products[1].Price},
		{OrderID: order.ID, ProductID: products[2].ID, Quantity: 1, Price: products[2].Price},
	}
	for i := range orderItems {
		db.Create(&orderItems[i])
	}

	// Retrieve the order with all its details
	var completeOrder Order
	db.Preload("Items.Product.Category").First(&completeOrder, order.ID)

	fmt.Printf("Order: %v for User ID: %v\n", completeOrder.OrderNumber, completeOrder.UserID)
	fmt.Printf("Order has %v items:\n", len(completeOrder.Items))
	for i, item := range completeOrder.Items {
		fmt.Printf("  Item %d: %v (Category: %v), Quantity: %v, Price: $%.2f\n",
			i+1,
			item.Product.Name,
			item.Product.Category.Name,
			item.Quantity,
			item.Price,
		)
	}

	// Retrieve a user's orders with all their details
	var userWithOrders User
	db.Preload("Orders.Items.Product.Category").First(&userWithOrders, user.ID)
	fmt.Printf("User: %v has %v orders\n", userWithOrders.Name, len(userWithOrders.Orders))
}
