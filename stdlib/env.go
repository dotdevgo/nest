package utils

import (
	"os"

	utils "github.com/Masterminds/goutils"
)

// GetEnv godoc
func GetEnv(name string, def string) string {
	val := os.Getenv(name)
	if utils.IsEmpty(val) {
		return def
	}
	return val
}
