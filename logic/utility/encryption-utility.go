package utility

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

var secretKey []byte = []byte("0123456789abcdef")

func CalculateHmacString(input string) string {
	mac := hmac.New(sha256.New, secretKey)
	mac.Write([]byte(input))
	return hex.EncodeToString(mac.Sum(nil))
}
