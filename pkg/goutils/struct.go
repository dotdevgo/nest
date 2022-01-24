package goutils

import (
	"github.com/gotidy/copy"
)

// Copy godoc
func Copy(target, source interface{}) {
	copiers := copy.New()
	copiers.Copy(target, source)
}
