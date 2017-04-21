package main

import (
	"encoding/hex"
)

func createCandidates(inp []byte) [255]string {
	var candidates [255]string
	for i := 0; i < 255; i++ {
		var candidate string
		key := byte(i + 1)
		for j := 0; j < len(inp); j++ {
			candidate += string(inp[j] ^ key)
		}
		candidates[i] = candidate
	}
	return candidates
}

func EvaluateCandidates(candidates [255]string) (int, string) {
	bestScore := 0.0
	var bestCandidate string
	var bestKey int
	for i := 0; i < len(candidates); i++ {
		score := ScoreString(candidates[i])
		if score > bestScore {
			bestScore = score
			bestCandidate = candidates[i]
			bestKey = i + 1
		}
	}
	return bestKey, bestCandidate
}

// This function expects a hex-encoded string
func SingleByteXOR(inp string) (int, string) {
	bytes, err := hex.DecodeString(inp)
	Check(err)
	candidates := createCandidates(bytes)
	key, answer := EvaluateCandidates(candidates)
	//return fmt.Sprintf("The key is %d with a value of \"%s\"", key, answer)
	return key, answer
}
