package main

import (
	"fmt"
)

func NewHeap() *heap {
	return &heap{arr: []int{0}}
}

type heap struct {
	arr []int
}

func (h *heap) add(n int) {
	h.arr = append(h.arr, n)
	cur := len(h.arr) - 1
	for cur > 1 {
		parent := cur / 2
		if h.arr[parent] > h.arr[cur] {
			h.arr[parent], h.arr[cur] = h.arr[cur], h.arr[parent]
			cur = parent
		} else {
			break
		}
	}
}

func (h *heap) del() int {
	res := h.arr[1]
	h.arr[1] = h.arr[len(h.arr)-1]
	cur := 1
	h.arr = h.arr[0 : len(h.arr)-1]
	for cur*2 < len(h.arr) {
		t := cur
		if h.arr[cur*2] < h.arr[cur] {
			h.arr[cur*2], h.arr[cur] = h.arr[cur], h.arr[cur*2]
			t = cur * 2
		}
		if cur*2+1 < len(h.arr) && h.arr[cur*2+1] < h.arr[cur] {
			h.arr[cur*2+1], h.arr[cur] = h.arr[cur], h.arr[cur*2+1]
			t = cur*2 + 1
		}
		if t == cur {
			break
		}else {
			cur = t
		}
	}
	return res
}

func main() {
	h := NewHeap()
	h.add(5)
	h.add(4)
	h.add(3)
	h.add(2)
	h.add(1)


	fmt.Println(h.del())
	fmt.Println(h.del())
	fmt.Println(h.del())
	fmt.Println(h.del())
	fmt.Println(h.del())
}
