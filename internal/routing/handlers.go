package routing

import (
	"github.com/labstack/echo/v4"
	"github.com/yus-works/capablanca/internal/templates"
	"go.uber.org/zap"
)


func RegisterRoutes(e *echo.Echo, log *zap.Logger) {
	e.GET("/", func(c echo.Context) error {
		log.Debug("RENDER: root endpoint")

		tables := []string{ "User", "Product", "Order", "OrderItem", "Review", "Category" }

		content := templates.Table(tables)

		return HTML(c, templates.Base("my title", templates.Bar(), content), 200)
	})
}
