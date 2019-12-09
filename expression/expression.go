package main

import (
	"fmt"
)

// https://wawlian.iteye.com/blog/1266397

var op = map[rune]bool{'+': true, '-': true, '*': true, '/': true}

type stack struct {
	arr []interface{}
}

func (s *stack) push(x interface{}) {
	s.arr = append(s.arr, x)
}

func (s *stack) top() interface{} {
	return s.arr[len(s.arr)-1]
}

func (s *stack) pop() interface{} {
	res := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return res
}

func toSuffix(s string) string {
	st := &stack{}
	arr := []rune(s)
	for _, r := range arr {
	}
}

func compute(suffix string) int {
	suffixArr := []rune(suffix)
	numberStack := &stack{}

	for _, r := range suffixArr {
		if _, ok := op[r]; ok {
			val1 := numberStack.pop().(int)
			val2 := numberStack.pop().(int)
			var tmp int
			switch r {
			case '+':
				tmp = val1 + val2
			case '-':
				tmp = val1 - val2
			case '*':
				tmp = val1 * val2
			case '/':
				tmp = val1 / val2
			}
			numberStack.push(tmp)
		} else {
			numberStack.push(int(r - '0'))
		}
	}

	return numberStack.pop().(int)
}

func main() {
	fmt.Println(compute("12+3*"))
	fmt.Println(compute("123*+"))
}
