package main

import (
	"crypto/sha1"
	"encoding/hex"
)

func HashPassword(pass string) string {
	hasher := sha1.New()
	hasher.Write([]byte(pass))
	return hex.EncodeToString(hasher.Sum(nil))
}
