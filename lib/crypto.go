package lib

import (
	"crypto/rand"
	"encoding/hex"
)

/**
 * From https://sosedoff.com/2014/12/15/generate-random-hex-string-in-go.html
 */
func RandomHex(length int) (string, error) {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}
