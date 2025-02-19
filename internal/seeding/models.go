package seeding

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null;size:100"`
	Email     string `gorm:"unique;not null;size:150"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Orders    []Order `gorm:"foreignKey:UserID"`
}

// Product represents a product in the catalog
type Product struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"unique;not null;size:200"`
	Description string  `gorm:"size:500"`
	Price       float64 `gorm:"not null"`
	Stock       int     `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Order represents an order made by a user
type Order struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `gorm:"not null"`
	Total     float64 `gorm:"not null"`
	Status    string  `gorm:"size:50;default:'pending'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Items     []OrderItem `gorm:"foreignKey:OrderID"`
}

// OrderItem represents an item in an order
type OrderItem struct {
	ID        uint    `gorm:"primaryKey"`
	OrderID   uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	Price     float64 `gorm:"not null"`
}

// Review represents a product review
type Review struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	ProductID uint   `gorm:"not null"`
	Rating    int    `gorm:"not null"`
	Comment   string `gorm:"size:500"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Category represents a product category
type Category struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique;not null;size:100"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Products  []Product `gorm:"many2many:category_products;"`
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&Product{},
		&Order{},
		&OrderItem{},
		&Review{},
		&Category{},
	)
}
