package logger

import (
	"hackerNewsApi/internal/components/config"
	"sync"

	"github.com/sirupsen/logrus"
)

type logrusWriter struct {
	Logger *logrus.Logger
}

type Logger interface {
	Printf(message string, args ...interface{})
}

var (
	once           sync.Once
	loggerInstance *logrusWriter
)

func NewLogger(config *config.Config) Logger {
	once.Do(func() {
		log := logrus.New()

		log.SetLevel(logrus.Level(config.LogLevel))
		log.SetFormatter(&logrus.JSONFormatter{})
		loggerInstance = &logrusWriter{Logger: log}
	})

	return loggerInstance
}

func (l *logrusWriter) Printf(message string, args ...interface{}) {
	l.Logger.Tracef(message, args...)
}
