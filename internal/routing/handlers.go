package routing

import (
	"github.com/labstack/echo/v4"
	"github.com/yus-works/capablanca/internal/repository"
	"github.com/yus-works/capablanca/internal/templates"
	"github.com/yus-works/capablanca/internal/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, log *zap.Logger, db *gorm.DB) {

	e.GET("/", func(c echo.Context) error {
		return c.HTML(200, "Hello, World!")
	})

	e.GET("/table/:name", func(c echo.Context) error {
		log.Debug("RENDER: root endpoint")

		namesSnake, err := repository.GetTableNames(db)
		if err != nil {
			log.Error("Failed to get table names")
		}

		namesPascal := make([]string, len(namesSnake))
		for i, name := range namesSnake {
			if len(name) == 0 {
				continue
			}

			namesPascal[i] = utils.SnakeToPascal(name)
		}

		tableName := c.Param("name")

		table, err := repository.GetTable(db, tableName)
		if err != nil {
			log.Error("Failed to get table data and metadata")
		}

		content := templates.Table(namesPascal, table)

		return HTML(c, templates.Base("my title", content), 200)
	})
}
