package main

import "fmt"

type Elem struct {
	low, high int
}

func qssort(n int, less func(i, j int) bool, swap func(i, j int)) {
	stack := []Elem{{0, n - 1}}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		low, high := top.low, top.high
		if low < high {
			pivotIndex := low + (high-low)/2
			swap(pivotIndex, high)
			storeIndex := low
			for i := low; i < high; i++ {
				if less(i, high) {
					swap(i, storeIndex)
					storeIndex++
				}
			}
			swap(storeIndex, high)
			stack = append(stack, Elem{low, storeIndex - 1})
			stack = append(stack, Elem{storeIndex + 1, high})
		}
	}
}

func main() {
	numbers := []int{1, 23, 54, 23, 78, 89}
	qssort(len(numbers),
		func(i, j int) bool { return numbers[i] < numbers[j] },
		func(i, j int) { numbers[i], numbers[j] = numbers[j], numbers[i] },
	)
	fmt.Println(numbers)
}
