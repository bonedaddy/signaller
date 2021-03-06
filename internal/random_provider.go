package internal

import (
	"crypto/rand"
	"fmt"
)

// RandomString returns new generated token with specified length
func RandomString(size int) string {
	b := make([]byte, size)
	rand.Read(b) // TODO: check may be can be error

	return fmt.Sprintf("%x", b)
}
