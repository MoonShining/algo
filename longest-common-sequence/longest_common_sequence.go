package main

// https://blog.csdn.net/hrn1216/article/details/51534607
// 算法导论
import (
	"fmt"
)

func LCS(a, b string) string {
	var seq []rune

	a1 := []rune(a)
	b1 := []rune(b)
	dp := make([][]int, len(b1)+1)
	for i := range dp {
		dp[i] = make([]int, len(a1)+1)
	}

	for i := 1; i < len(b1); i++ {
		for j := 1; j < len(a1); j++ {
			if a1[i] == b1[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	i, j := len(b1)-1, len(a1)-1
	for i >= 0 && j >= 0 {
		if b1[i] == a1[j] {
			seq = append(seq, b1[i])
			j--
			i--
		} else {
			if dp[i][j-1] > dp[i-1][j] {
				j--
			} else {
				i--
			}
		}
	}
	return reverse(seq)
}

func reverse(a []rune) string {
	i, j := 0, len(a)-1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
	return string(a)
}

func main() {
	lcs := LCS("abcfbc", "abfcab")
	fmt.Println(lcs)
}
