package seeding

import (
	"time"

	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) error {
	if err := nuke(db); err != nil {
		return err
	}

	if err := migrate(db); err != nil {
		return err
	}

	return seed(db)
}

func seed(db *gorm.DB) error {
	// Users
	users := []User{
		{Username: "john_doe", Email: "john@example.com", Password: "hashedpassword", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Username: "jane_doe", Email: "jane@example.com", Password: "hashedpassword", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	if err := db.Create(&users).Error; err != nil {
		return err
	}

	// Products
	products := []Product{
		{Name: "Laptop", Description: "High-end gaming laptop", Price: 1500.99, Stock: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Smartphone", Description: "Latest smartphone model", Price: 999.99, Stock: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	if err := db.Create(&products).Error; err != nil {
		return err
	}

	// Orders
	orders := []Order{
		{UserID: 1, Total: 1500.99, Status: "pending", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserID: 2, Total: 999.99, Status: "shipped", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	if err := db.Create(&orders).Error; err != nil {
		return err
	}

	// OrderItems
	orderItems := []OrderItem{
		{OrderID: 1, ProductID: 1, Quantity: 1, Price: 1500.99},
		{OrderID: 2, ProductID: 2, Quantity: 1, Price: 999.99},
	}
	if err := db.Create(&orderItems).Error; err != nil {
		return err
	}

	// Reviews
	reviews := []Review{
		{UserID: 1, ProductID: 1, Rating: 5, Comment: "Great product!", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{UserID: 2, ProductID: 2, Rating: 4, Comment: "Good but expensive.", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	if err := db.Create(&reviews).Error; err != nil {
		return err
	}

	// Categories
	categories := []Category{
		{Name: "Electronics", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{Name: "Accessories", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	if err := db.Create(&categories).Error; err != nil {
		return err
	}

	return nil
}

func nuke(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&User{}, &Product{}, &Order{}, 
		&OrderItem{}, &Review{}, &Category{},
	)
}
