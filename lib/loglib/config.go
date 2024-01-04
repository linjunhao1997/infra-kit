package loglib

import (
	"log"
	"log/slog"
	"os"
	"path/filepath"
)

type Config struct {
	Level    string
	FilePath string
}

type Option func(c *Config)

func Init(opts ...Option) {
	var config Config
	for _, f := range opts {
		f(&config)
	}
	var err error
	var logger *slog.Logger
	defer func() {
		if err != nil {
			log.Printf("failed init logger: %v\n", err)
			logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: caseLevel(config.Level)}))
		}
		slog.SetDefault(logger)
	}()
	if len(config.FilePath) > 0 {
		err = os.MkdirAll(filepath.Dir(config.FilePath), 0666)
		if err != nil {
			return
		}

		file, err := os.Create(config.FilePath)
		if err != nil {
			return
		}
		defer file.Close()
		logger = slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{Level: caseLevel(config.Level)}))
	} else {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: caseLevel(config.Level)}))
	}
}

func caseLevel(level string) slog.Level {
	switch level {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
