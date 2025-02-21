package main

import (
	"io"

	"github.com/labstack/echo/v4"
	"github.com/yus-works/capablanca/internal/repository"
	"github.com/yus-works/capablanca/internal/logging"
	"github.com/yus-works/capablanca/internal/middleware"
	"github.com/yus-works/capablanca/internal/routing"

	"go.uber.org/zap"
)

func main() {
	logger := logging.SetupLogger()
	defer logger.Sync()

	logger.Debug("Starting application...")

	db := repository.SetupDb(logger)

	e := echo.New()

	// Disable Echo's built-in logger output.
	e.Logger.SetOutput(io.Discard)

	middleware.RegisterMiddleware(e, logger)

	routing.RegisterRoutes(e, logger, db)

	// Start the server.
	logger.Info("Starting server", zap.String("addr", ":8080"))
	if err := e.Start(":8080"); err != nil {
		logger.Fatal("server encountered a fatal error", zap.Error(err))
	}
}
