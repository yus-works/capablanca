package routing

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// renders a template and responds to the request
func HTML(c echo.Context, cmp templ.Component, status int) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	c.Response().WriteHeader(status)
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
