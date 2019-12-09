package main

import "log"

func isValid(s string) bool {
	runes := []rune(s)

	stack := []rune{}

	for _, r := range runes {
		if r == '{' || r == '[' || r == '(' {
			stack = append(stack, r)
		} else if r == '}' || r == ']' || r == ')' {
			if len(stack) == 0 {
				return false
			}
			topVal := stack[len(stack)-1]

			if r == '}' && topVal == '{' {
				stack = stack[:len(stack)-1]
			} else if r == ']' && topVal == '[' {
				stack = stack[:len(stack)-1]
			} else if r == ')' && topVal == '(' {
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
	}
	return len(stack) == 0
}

func main() {
	res := isValid("()[]{}")
	log.Print(res)
}
