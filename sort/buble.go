package main

import "fmt"

func Bubble(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j+1], a[j] = a[j], a[j+1]
			}
		}
	}
}

func Quick(a []int, lo, hi int) {
	if lo < hi {
		j := partition(a, lo, hi)
		Quick(a, lo, j-1)
		Quick(a, j+1, hi)
	}
}

func partition(a []int, lo, hi int) int {
	i := lo
	pivot := a[hi]
	for j := lo; j <= hi; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	return i
}

func Select(a []int) {
	for i := 0; i < len(a); i++ {
		min := i
		for j := i; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}


func main() {
	a := []int{3, 9, 1, 4, 5, 6, 4, 7, 1}

	//Bubble(a)
	//Quick(a, 0, len(a)-1)
	Select(a)
	fmt.Println(a)
}
