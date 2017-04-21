package main

import (
	"crypto/aes"
)

func aesECB(file, key string) string {
	size := len(key)
	input := ReadFile(file)
	output := make([]byte, len(input))
	cipher, error := aes.NewCipher([]byte(key))
	Check(error)
	for i, j := 0, size; i < len(input); i, j = i+size, j+size {
		cipher.Decrypt(output[i:j], []byte(input)[i:j])
	}
	return string(output)
}
