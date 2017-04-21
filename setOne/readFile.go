package main

import (
	"encoding/base64"
	"io/ioutil"
	"strings"
)

// reads a base64-encoded file
// and output a single string
func ReadFile(filepath string) string {
	dat, err := ioutil.ReadFile(filepath)
	Check(err)

	sanitizedString := strings.Replace(string(dat), "\n", "", -1)

	output, errTwo := base64.StdEncoding.DecodeString(sanitizedString)
	Check(errTwo)

	return string(output)
}

// reads a file line by line,
// appending each line to an array
func ReadLines(filepath string) []string {
	dat, err := ioutil.ReadFile(filepath)
	Check(err)

	lines := strings.Split(string(dat), "\n")
	return lines
}
