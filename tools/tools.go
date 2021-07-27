package tools

import (
	"crypto/sha256"
	"encoding/hex"
)

func EncodingSha256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	sum := h.Sum(nil)
	s := hex.EncodeToString(sum)
	return string(s)
}
