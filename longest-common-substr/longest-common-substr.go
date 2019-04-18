package main

import "fmt"

//https://blog.csdn.net/u010397369/article/details/38979077
//http://www.ahathinking.com/archives/122.html

func LongestCommonSubstr(a, b string) (int, int) {
	var length int
	var index int

	a1, b1 := []rune(a), []rune(b)
	dp := make([][]int, len(b1))
	for i := range dp {
		dp[i] = make([]int, len(a1))
	}
	for i := 0; i < len(b1); i++ {
		for j := 0; j < len(a1); j++ {
			if a1[j] == b1[i] {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
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
