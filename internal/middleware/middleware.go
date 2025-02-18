package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yus-works/capablanca/internal/logging"
	"go.uber.org/zap"
)


func RegisterMiddleware(e *echo.Echo, l *zap.Logger) {
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			l.Error("PANIC recovered", zap.Error(err), zap.String("stack", string(stack)))
			return err
		},
		StackSize: 1 << 10, // 1KB stack trace
	}))

	// Optional: add custom request logging middleware that uses Zap.
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			method := c.Request().Method
			uri := c.Request().RequestURI

			// The coloredMethod is part of the main message so the terminal will render the colors.
			l.Debug(logging.ColorMethod(method)+" "+uri,
				zap.String("method", method), // raw method for structured logging
				zap.String("uri", uri),
			)
			return next(c)
		}
	})

	// TODO: maybe add response code to logging too?
}
