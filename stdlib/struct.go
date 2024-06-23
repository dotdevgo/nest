package utils

import (
	"github.com/gotidy/copy"
)

var copiers = copy.New()

// Copy godoc
func Copy(target, source interface{}) {
	// copiers := copy.New()
	copiers.Copy(target, source)
}
