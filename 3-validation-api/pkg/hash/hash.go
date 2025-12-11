package hash

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateHash(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
