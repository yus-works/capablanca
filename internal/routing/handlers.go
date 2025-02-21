package routing

import (
	"github.com/labstack/echo/v4"
	"github.com/yus-works/capablanca/internal/repository"
	"github.com/yus-works/capablanca/internal/templates"
	"go.uber.org/zap"
	"gorm.io/gorm"
)


func RegisterRoutes(e *echo.Echo, log *zap.Logger, db *gorm.DB) {
	e.GET("/", func(c echo.Context) error {
		log.Debug("RENDER: root endpoint")

		names, err := repository.GetTableNames(db)
		if err != nil {
			log.Error("Failed to get table names")
		}

		content := templates.Table(names)

		return HTML(c, templates.Base("my title", templates.Bar(), content), 200)
	})
}
