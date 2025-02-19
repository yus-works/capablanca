package db

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Example model
type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100"`
}

func SetupDb(l *zap.Logger) {
	// Connect to database using GORM ---
	dsn := "capablanca:secret@tcp(127.0.0.1:3306)/capablanca?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		l.Fatal("failed to connect to database", zap.Error(err))
	}

	// Auto-migrate the schema.
	if err := db.AutoMigrate(&User{}); err != nil {
		l.Fatal("failed to auto-migrate schema", zap.Error(err))
	}
}
