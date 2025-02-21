package repository

import (
	"github.com/yus-works/capablanca/internal/seeding"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDb(l *zap.Logger) *gorm.DB {
	// Connect to database using GORM ---
	dsn := "capablanca:secret@tcp(127.0.0.1:3306)/capablanca?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		l.Fatal("failed to connect to database", zap.Error(err))
	}

	if err := seeding.SeedDatabase(db); err != nil {
		l.Fatal("failed to seed database", zap.Error(err))
	}

	return db
}
