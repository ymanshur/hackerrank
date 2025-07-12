package main

import "fmt"

type Queue[T any] []T

func (q *Queue[T]) PushBack(val T) {
	*q = append(*q, val)
}

func (q *Queue[T]) PopFront() (T, bool) {
	var zero T
	if len(*q) == 0 {
		return zero, false
	}

	val := (*q)[0]
	(*q)[0] = zero
	*q = (*q)[1:]
	return val, true
}

func (q *Queue[T]) Front() (T, bool) {
	var zero T
	if len(*q) == 0 {
		return zero, false
	}
	return (*q)[0], true
}

func main() {
	var queue Queue[int32]
	// val int32

	// queue.PushBack(43)
	// queue.PopFront()
	// queue.PushBack(14)
	// val, _ = queue.Front()
	// fmt.Println(val)
	// queue.PushBack(28)
	// val, _ = queue.Front()
	// fmt.Println(val)
	// queue.PushBack(60)
	// queue.PushBack(78)
	// queue.PopFront()
	// queue.PopFront()

	var q int32
	fmt.Scan(&q)

	for i := int32(0); i < q; i++ {
		var t, x int32
		fmt.Scan(&t)

		switch t {
		case 1:
			fmt.Scan(&x)
			queue.PushBack(x)
		case 2:
			queue.PopFront()
		case 3:
			val, _ := queue.Front()
			fmt.Println(val)
		}
	}
}
