package helper

import (
	"math/rand"
	"time"
)

// random from timenano
var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// whitelisted char to generate random string
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Generate random string with specific length
func GenerateRandStr(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
