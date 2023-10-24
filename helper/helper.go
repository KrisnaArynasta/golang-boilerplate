package Helper

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func HmacSHA256(data, secretKey string) string {
	key := []byte(secretKey)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
