package backend

import (
	"math/rand"
	"time"
)

// This code was sourced from https://www.calhoun.io/creating-random-strings-in-go/
// This code was written by Jon Calhoun

// Characters available
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Function to create a random seed
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// Function to generate a sting of given length made from provided characteres
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// Random string generator
func String(length int) string {
	return StringWithCharset(length, charset)
}
