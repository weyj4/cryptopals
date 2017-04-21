package main

func decodeLines(lines []string) ([]int, []string) {
	decodedLines := make([]string, len(lines))
	keys := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		key, ans := SingleByteXOR(lines[i])
		decodedLines[i] = ans
		keys[i] = key
	}
	return keys, decodedLines
}

func evaluateLines(lines []string) (int, string) {
	bestScore := 0.0
	var bestLine string
	var index int
	for i := 0; i < len(lines); i++ {
		score := ScoreString(lines[i])
		if score > bestScore {
			bestScore = score
			bestLine = lines[i]
			index = i
		}
	}
	return index, bestLine
}

func DetectSingleByteXOR(filepath string) (int, string) {
	lines := ReadLines(filepath)
	keys, answers := decodeLines(lines)
	idx, answer := evaluateLines(answers)
	key := keys[idx]
	return key, answer
}
