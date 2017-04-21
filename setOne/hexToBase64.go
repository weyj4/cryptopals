package main

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(str string) string {
	bytes, err := hex.DecodeString(str)
	Check(err)
	return base64.StdEncoding.EncodeToString(bytes)
}
