package routing

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/yus-works/capablanca/internal/repository"
	"github.com/yus-works/capablanca/internal/templates"
	"go.uber.org/zap"
	"gorm.io/gorm"
)


func RegisterRoutes(e *echo.Echo, log *zap.Logger, db *gorm.DB) {
	e.GET("/table/:name", func(c echo.Context) error {
		log.Debug("RENDER: root endpoint")

		names, err := repository.GetTableNames(db)
		if err != nil {
			log.Error("Failed to get table names")
		}


		var namesClean []string
		for _, name := range names {
			if strings.Contains(name, "_") { continue }

			if len(name) > 0 {
				title := strings.ToUpper(name[:1]) + name[1:]
				namesClean = append(namesClean, title)
			}
		}

		tableName := c.Param("name")

		columns, err := repository.GetTableColumns(db, tableName)
		if err != nil {
			log.Error("Failed to get column names for table ")
		}

		tableData, err := repository.GetTableData(db, tableName)
		if err != nil {
			log.Error("Failed to get column names for table ")
		}

		content := templates.Table(tableName, namesClean, columns, tableData)

		return HTML(c, templates.Base("my title", content), 200)
	})
}
