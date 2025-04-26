package logger

import (
	"github.com/mixdone/fly-api/internal/config"

	"github.com/sirupsen/logrus"
)

func NewLogger(cfg *config.Config) *logrus.Logger {
	log := logrus.New()

	level, err := logrus.ParseLevel(cfg.Log.Level)
	if err != nil {
		log.Fatalf("Invalid log level: %v", err)
	}
	log.SetLevel(level)

	if cfg.Log.Format == "json" {
		log.SetFormatter(&logrus.JSONFormatter{})
	} else {
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	return log
}
