package main

import (
	"fmt"
)

func substr(s, t string) int {
	ss := []rune(s)
	tt := []rune(t)

	res := -1
	for i := 0; i <= len(ss)-len(tt); i++ {
		var j int
		for j = 0; j < len(tt); j++ {
			if ss[i+j] != tt[j] {
				break
			}
		}

		if j == len(tt) {
			res = i
			break
		}
	}
	return res
}
func main() {
	n := substr("how are you", "are")
	fmt.Print(n)
}
