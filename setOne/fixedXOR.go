package main

import (
	"encoding/hex"
)

// This function expects hex strings as input
func FixedXOR(str1, str2 string) string {
	unmatched := len(str1) != len(str2)
	CheckCustom(unmatched, "String lengths don't match")

	bytes1, err1 := hex.DecodeString(str1)
	Check(err1)
	bytes2, err2 := hex.DecodeString(str2)
	Check(err2)

	var outputBytes = make([]byte, len(bytes1))

	for i := 0; i < len(bytes1); i++ {
		outputBytes[i] = (bytes1[i] ^ bytes2[i])
	}

	return hex.EncodeToString(outputBytes)
}
