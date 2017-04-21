package main

import (
	"strings"
)

func scoreChar(c byte) float64 {
	switch c {
	case ' ':
		return 15.000
	case 'e':
		return 12.702
	case 't':
		return 9.056
	case 'a':
		return 8.167
	case 'o':
		return 7.507
	case 'i':
		return 6.996
	case 'n':
		return 6.749
	case 's':
		return 6.327
	case 'h':
		return 6.094
	case 'r':
		return 5.987
	case 'd':
		return 4.253
	case 'l':
		return 4.025
	case 'c':
		return 2.782
	case 'u':
		return 2.758
	case 'm':
		return 2.406
	case 'w':
		return 2.360
	case 'f':
		return 2.228
	case 'g':
		return 2.015
	case 'y':
		return 1.974
	case 'p':
		return 1.974
	case 'b':
		return 1.492
	case 'v':
		return 0.978
	case 'k':
		return 0.772
	case 'j':
		return 0.153
	case 'x':
		return 0.150
	case 'q':
		return 0.095
	case 'z':
		return 0.074
	default:
		return 0
	}
}

func ScoreString(inp string) float64 {
	inp = strings.ToLower(inp)
	score := 0.0
	for i := 0; i < len(inp); i++ {
		score += scoreChar(inp[i])
	}
	return score / float64(len(inp))
}
