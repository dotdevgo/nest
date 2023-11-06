package logger

import (
	"fmt"
	"os"

	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

// Log
func Log(format string, v ...any) {
	logrus.Info(fmt.Sprintf(format, v...))
}

// Info godoc
func Info(v ...any) {
	logrus.Info(v...)
}

// Error godoc
func Error(v ...any) {
	log.Error(v...)
}

// Warn godoc
func Warn(v ...any) {
	log.Warn(v...)
}

// Fatal godoc
func Fatal(v ...any) {
	log.Fatal(v...)
}

// Panic godoc
func Panic(v ...any) {
	log.Panic(v...)
}

// PanicOnError godoc
func PanicOnError(err error) {
	if err != nil {
		Panic(err)
	}
}

// FatalOnError godoc
func FatalOnError(err error) {
	if err != nil {
		Panic(err)
	}
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
