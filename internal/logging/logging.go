package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SetupLogger() *zap.Logger {
	// --- Setup Zap logger with two separate cores ---

	// Console core: pretty print with colors.
	// We want to show only debug (and error) messages on the console.
	consoleEncoderConfig := zap.NewDevelopmentEncoderConfig()
	consoleEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // adds colors
	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(consoleEncoderConfig),
		zapcore.Lock(os.Stdout),
		// Only enable Debug-level (and errors) on console.
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zap.DebugLevel || lvl >= zap.ErrorLevel
		}),
	)

	// File core: structured JSON output.
	fileEncoderConfig := zap.NewProductionEncoderConfig()
	logFile, err := os.OpenFile("structured.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic("failed to open log file: " + err.Error())
	}
	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(fileEncoderConfig),
		zapcore.AddSync(logFile),
		zap.DebugLevel, // all logs (debug and above) go to file
	)

	// Combine the two cores.
	combinedCore := zapcore.NewTee(
		consoleCore, // writes only debug (and errors) to stdout
		fileCore,    // writes structured JSON to file
	)

	return zap.New(combinedCore, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
}

func ColorMethod(method string) string {
	var coloredMethod string
	switch method {
	case "GET":
		// Green background with black text.
		coloredMethod = "\033[42;30m" + method + "\033[0m"
	case "POST":
		// Blue background with white text.
		coloredMethod = "\033[44;97m" + method + "\033[0m"
	case "PUT":
		// Magenta background with white text.
		coloredMethod = "\033[45;97m" + method + "\033[0m"
	case "DELETE":
		// Red background with white text.
		coloredMethod = "\033[41;97m" + method + "\033[0m"
	default:
		coloredMethod = method
	}

	return coloredMethod
}
