package utils

import (
	"os"

	utils "github.com/Masterminds/goutils"
)

// Getenv godoc
func Getenv(name string, def string) string {
	val := os.Getenv(name)
	if utils.IsEmpty(val) {
		return def
	}
	return val
}
