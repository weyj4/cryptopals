package main

func Padding(block string, length int) string {
	bytes := []byte(block)
	pad := byte(4)
	numBytes := length - len(block)
	for i := 0; i < numBytes; i++ {
		bytes = append(bytes, pad)
	}
	return string(bytes)
}
