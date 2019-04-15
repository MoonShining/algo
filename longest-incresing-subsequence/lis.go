package main

import (
	"fmt"
)

func lis(a []int) int {
	tmp := make([]int, len(a))

	res := 1

	for i := 0; i < len(a); i++ {
		tmp[i] = 1

		for j := 0; j < i; j++ {
			if a[j] <= a[i] && tmp[j]+1 > tmp[i] {
				tmp[i] = tmp[j] + 1
			}
		}

		if tmp[i] > res {
			res = tmp[i]
		}
	}
	return res
}

func main() {
	fmt.Println(lis([]int{5, 3, 4, 8, 6, 7}))
}
