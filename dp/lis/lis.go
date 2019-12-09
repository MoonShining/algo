package main

import (
	"fmt"
)

func lis(a []int) int {
	max := 0
	dp := make([]int, len(a))
	for i, _ := range dp {
		dp[i] = 1
	}

	for i := 1; i < len(a); i++ {
		for j := 0; j < i; j++ {
			if a[i] > a[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}

	for _, v := range dp {
		if max < v {
			max = v
		}
	}

	return max
}

func main() {
	fmt.Println(lis([]int{5, 3, 4, 8, 6, 7}))
}
