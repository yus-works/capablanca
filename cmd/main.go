package main

import (
	"io"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yus-works/capablanca/internal/logging"
	"github.com/yus-works/capablanca/internal/templates"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go.uber.org/zap"
)

// Example model
type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100"`
}

// renders a template and responds to the request
func HTML(c echo.Context, cmp templ.Component, status int) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	c.Response().WriteHeader(status)
	return cmp.Render(c.Request().Context(), c.Response().Writer)
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
	// Remove Echo’s default middleware.Logger() to prevent it from logging JSON to stdout.
	e.Use(middleware.Recover())

	// Optional: add custom request logging middleware that uses Zap.
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			method := c.Request().Method
			uri := c.Request().RequestURI

			// The coloredMethod is part of the main message so the terminal will render the colors.
				logger.Debug("REQUEST: " + logging.ColorMethod(method) + " " + uri,
				zap.String("method", method), // raw method for structured logging
				zap.String("uri", uri),
			)
			return next(c)
		}
	})

	// Example route.
	e.GET("/", func(c echo.Context) error {
		logger.Debug("RENDER: root endpoint")
		return HTML(c, templates.Base("joe"), 200)
	})

	// Start the server.
	logger.Info("Starting server", zap.String("addr", ":8080"))
	if err := e.Start(":8080"); err != nil {
		logger.Fatal("server encountered a fatal error", zap.Error(err))
	}
}
