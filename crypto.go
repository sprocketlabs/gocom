package gocom

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func SHA256String(s string) string {
	b := []byte(s)
	outArr := sha256.Sum256(b)
	outSlice := outArr[:]
	return hex.EncodeToString(outSlice)
}

func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64Decoe(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
