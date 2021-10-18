package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func Initialize() *logrus.Logger {
	return &logrus.Logger{
		Out:       os.Stderr,
		Level:     logrus.DebugLevel,
		Formatter: &logrus.JSONFormatter{},
	}
}
