type pq struct {
	arr []int
}

func (q *pq) push(n int) {
	q.arr = append(q.arr, n)
	i := len(q.arr) - 1
	for i > 0 {
		p := (i - 1) / 2
		if q.arr[p] > q.arr[i] {
			q.arr[p], q.arr[i] = q.arr[i], q.arr[p]
		}
		i = p
	}
}

func (q *pq) poll() int {
	ret := q.arr[0]
	q.arr[0], q.arr[len(q.arr)-1] = q.arr[len(q.arr)-1], q.arr[0]
	q.arr = q.arr[:len(q.arr)-1]

	i := 0
	p := 0
	for {
		p = i*2 + 1
		if p >= len(q.arr) {
			break
		}

		if p < len(q.arr) && p+1 < len(q.arr) {
			if q.arr[p] > q.arr[p+1] {
				p = p + 1
			}
		}
		if q.arr[p] < q.arr[i] {
			q.arr[p], q.arr[i] = q.arr[i], q.arr[p]
			i = p
		} else {
			break
		}
	}
	return ret
}

