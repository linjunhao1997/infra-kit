package loglib

import (
	"log"
	"log/slog"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
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
		r := &lumberjack.Logger{Filename: config.FilePath, LocalTime: true, MaxSize: 1, MaxAge: 3, MaxBackups: 5, Compress: true}

		logger = slog.New(slog.NewJSONHandler(r, &slog.HandlerOptions{Level: caseLevel(config.Level)}))
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
