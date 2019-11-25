package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}

	// dp
	a := 0
	b := 1
	c := 0
	for i := 2; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}

func main() {
	i := fib(5)
	fmt.Println(i)
}
