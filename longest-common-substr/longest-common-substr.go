package main

import "fmt"

//https://blog.csdn.net/u010397369/article/details/38979077
//http://www.ahathinking.com/archives/122.html

func LongestCommonSubstr(a, b string) (int, int) {
	if len(a) == 0 || len(b) == 0 {
		return 0, 0
	}

	var length int
	var index int

	a1, b1 := []rune(a), []rune(b)
	dp := make([][]int, len(b1))
	for i := range dp {
		dp[i] = make([]int, len(a1))
	}
	for i := 0; i < len(b1); i++ {
		if b1[i] == a1[0] {
			dp[i][0] = 1
		}
	}
	for i := 0; i < len(a1); i++ {
		if a1[i] == b1[0] {
			dp[0][i] = 1
		}
	}
	for i := 1; i < len(b1); i++ {
		for j := 1; j < len(a1); j++ {
			if a1[j] == b1[i] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > length {
					length = dp[i][j]
					index = j - length + 1
				}
			} else {
				dp[i][j] = 0
			}
		}
	}
	return length, index
}

func main() {
	a, b := "abcde", "bcde"
	length, index := LongestCommonSubstr(a, b)
	fmt.Println(a[index : index+length])
}
