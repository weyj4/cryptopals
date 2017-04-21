package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func hammingDistance(a, b string) int {
	score := 0

	for i := 0; i < len(a); i++ {
		xored := a[i] ^ b[i]
		score += strings.Count(fmt.Sprintf("%b", xored), "1")
	}

	return score
}

func getKeySize(input string, max int) int {
	for max*2 > len(input) {
		max = max - 1
	}
	bestDist := max + 1
	var bestLength int
	for i := 2; i < max; i++ {
		firstString := input[0:i]
		secondString := input[i : i*2]
		dist := int(hammingDistance(firstString, secondString) / i)
		if dist < bestDist {
			bestDist = dist
			bestLength = i
		}
	}
	return bestLength
}

func breakIntoChunks(input string, size int) []string {
	output := make([]string, size)
	for i, v := range input {
		output[i%size] += string(v)
	}
	return output
}

func hexEncodeChunks(chunkedInput []string) []string {
	output := make([]string, len(chunkedInput))
	for i, v := range chunkedInput {
		output[i] = string(hex.EncodeToString([]byte(v)))
	}
	return output
}

func decodeChunks(chunkedInput []string) []string {
	output := make([]string, len(chunkedInput))
	keys := make([]int, len(chunkedInput))
	for i, v := range chunkedInput {
		keys[i], output[i] = SingleByteXOR(v)
	}
	return output
}

func unChunk(chunkedInput []string) string {
	output := make([]string, len(chunkedInput[0]))
	for _, v := range chunkedInput {
		for i, k := range v {
			output[i] += string(k)
		}
	}
	return strings.Join(output, "")
}

func BreakRepeatingKey(filepath string) string {
	encoded := ReadFile(filepath)
	keySize := getKeySize(encoded, 40)
	fmt.Println(keySize)
	keySize = 29
	chunks := breakIntoChunks(encoded, keySize)
	// we need to work with hex data
	hexChunks := hexEncodeChunks(chunks)
	decodedChunks := decodeChunks(hexChunks)
	unchunked := unChunk(decodedChunks)
	return unchunked
}
