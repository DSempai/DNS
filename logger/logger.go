package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Initialize return new logging instance.
func Initialize() *logrus.Logger {
	return &logrus.Logger{
		Out:       os.Stderr,
		Level:     logrus.DebugLevel,
		Formatter: &logrus.JSONFormatter{},
	}
}
