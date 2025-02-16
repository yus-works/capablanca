package main

import (
	"log"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yus-works/capablanca/internal/templates"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Example model
type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100"`
}

func HTML(c echo.Context, cmp templ.Component, status int) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	c.Response().WriteHeader(status)
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func main() {
	// Connect to database using GORM
	dsn := "capablanca:secret@tcp(127.0.0.1:3306)/capablanca?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// NOTE: Auto-migrate schema
	db.AutoMigrate(&User{})

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Example route
	e.GET("/", func(c echo.Context) error {
		return HTML(c, templates.Base("joe"), 200)
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
