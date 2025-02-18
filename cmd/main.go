package main

import (
	"io"

	"github.com/labstack/echo/v4"
	"github.com/yus-works/capablanca/internal/logging"
	"github.com/yus-works/capablanca/internal/middleware"
	"github.com/yus-works/capablanca/internal/routing"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go.uber.org/zap"
)

// Example model
type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100"`
}

func main() {
	logger := logging.SetupLogger()
	defer logger.Sync()

	logger.Debug("Starting application...")

	// Connect to database using GORM ---
	dsn := "capablanca:secret@tcp(127.0.0.1:3306)/capablanca?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("failed to connect to database", zap.Error(err))
	}

	// Auto-migrate the schema.
	if err := db.AutoMigrate(&User{}); err != nil {
		logger.Fatal("failed to auto-migrate schema", zap.Error(err))
	}

	e := echo.New()

	// Disable Echo's built-in logger output.
	e.Logger.SetOutput(io.Discard)

	middleware.RegisterMiddleware(e, logger)

	routing.RegisterRoutes(e, logger)

	// Start the server.
	logger.Info("Starting server", zap.String("addr", ":8080"))
	if err := e.Start(":8080"); err != nil {
		logger.Fatal("server encountered a fatal error", zap.Error(err))
	}
}
