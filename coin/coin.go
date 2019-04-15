package main

import (
	"fmt"
)

var (
	coins = []int{1, 3, 5} // 面值为1、3、5的硬币若干
)

func leastCoin(sum int) int {
	if sum == 0 {
		return 0
	}

	// 初始化 全部用面值为1的
	tmp := make([]int, sum+1)
	for i, _ := range tmp {
		tmp[i] = i
	}

	// 每一个当前面额
	for i := 1; i <= sum; i++ {
		//当前面额>=可选硬币面额
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 && tmp[i-coins[j]]+1 < tmp[i] {
				tmp[i] = tmp[i-coins[j]] + 1
			}
		}
	}

	return tmp[sum]
}

func main() {
	fmt.Println(leastCoin(9))
}
