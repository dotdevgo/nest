package utils

import "github.com/dotdevgo/nest/pkg/logger"

// NoError godoc
func NoError(err error) {
	if err != nil {
		logger.Error(err)
	}
}

// NoErrorOrPanic godoc
func NoErrorOrPanic(err error) {
	if err != nil {
		logger.Panic(err)
	}
}

// NoErrorOrFatal godoc
func NoErrorOrFatal(err error) {
	if err != nil {
		logger.Panic(err)
	}
}
