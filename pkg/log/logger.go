package log

import (
	"log/slog"
	"os"
	"path/filepath"
)

type Logger interface {
	Error(msg string, keysAndValues ...any)
	Info(msg string, keysAndValues ...any)
	Warn(msg string, keysAndValues ...any)
}

func InitSlog(logToFile bool, filePath string) Logger {
	var output *os.File
	var err error

	if logToFile {
		dir := filepath.Dir(filePath)

		if _, err = os.Stat(dir); os.IsNotExist(err) {
			err = os.MkdirAll(dir, 0755)
			if err != nil {
				panic(err)
			}

		} else if err != nil {
			panic(err)
		}

		output, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
	} else {
		output = os.Stdout
	}

	handler := slog.NewJSONHandler(output, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	logger := slog.New(handler)
	slog.SetDefault(logger)

	return logger
}
