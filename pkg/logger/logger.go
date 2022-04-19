package logger

import (
	"os"

	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

// Log godoc
func Log(v ...any) {
	logrus.Info(v...)
}

// Fatal godoc
func Fatal(v ...any) {
	log.Fatal(v...)
}

// Panic godoc
func Panic(v ...any) {
	log.Panic(v...)
}

// Error godoc
func Error(v ...any) {
	log.Error(v...)
}

// Init godoc
func Init() {
	// Fatal as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	// logrus.SetLevel(logrus.WarnLevel)
}
