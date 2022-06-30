package utils

import (
	"github.com/gotidy/copy"
)

// Copy godoc
func Copy(target, source interface{}) {
	copiers := copy.New()
	copiers.Copy(target, source)
}

// // CopyMap godoc
// func CopyMap(target interface{}, source interface{}) {
// 	for index, element := range source {
// 		target[index] = element
// 	}
// }
