package main

import (
	"encoding/hex"
)

func getKeyByte(idx, length int) int {
	return idx % length
}

func RepeatingKeyXOR(key string, input string) string {
	acc := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		acc[i] = input[i] ^ key[getKeyByte(i, len(key))]
	}
	return hex.EncodeToString(acc)
}
