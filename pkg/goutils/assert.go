package goutils

import "github.com/labstack/gommon/log"

// NoError godoc
func NoError(err error) {
	if err != nil {
		log.Error(err)
	}
}

// NoErrorOrPanic godoc
func NoErrorOrPanic(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// NoErrorOrFatal godoc
func NoErrorOrFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
