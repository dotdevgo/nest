package utils

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var DefaultLength = 14

const charSet = "abcdedfghijklmnopqrstABCDEFGHIJKLMNOP"

// RandomString godoc
func RandomString(length *int) string {
	if nil == length {
		length = &DefaultLength
	}

	// TODO: fix deprecation
	rand.Seed(time.Now().Unix())

	var output strings.Builder
	for i := 0; i < *length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}

	return output.String()
}

// RandomToken godoc
func RandomToken() string {
	timestamp := time.Now().Unix()
	return fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprint(timestamp))))[:45]
}
