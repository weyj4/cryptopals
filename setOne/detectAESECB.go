package main

import (
	"encoding/hex"
)

func hasDuplicates(chunks []string) bool {
	found := make(map[string]bool)
	for _, v := range chunks {
		if !found[v] {
			found[v] = true
		} else {
			return true
		}
	}
	return false
}

func decodeHex(input []string) []string {
	output := make([]string, len(input))
	for i, v := range input {
		decoded, err := hex.DecodeString(v)
		Check(err)
		output[i] = string(decoded)
	}
	return output
}

func keySizeChunks(input string, keysize int) []string {
	output := make([]string, len(input)/keysize)
	for i, j, k := 0, 16, 0; j <= len(input); i, j, k = i+16, j+16, k+1 {
		output[k] = input[i:j]
	}
	return output
}

func DetectAESECB(filepath string) string {
	dat := ReadLines(filepath)
	decodedDat := decodeHex(dat)
	for i := 0; i < len(decodedDat); i++ {
		chunks := keySizeChunks(decodedDat[i], 16)
		if hasDuplicates(chunks) {
			return dat[i]
		}
	}
	return ""
}
