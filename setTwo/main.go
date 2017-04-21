package main

import "fmt"

func main() {
	ans := Padding("YELLOW SUBMARINE", 20)
	fmt.Printf("%s, %d\n", ans, len(ans))
}
