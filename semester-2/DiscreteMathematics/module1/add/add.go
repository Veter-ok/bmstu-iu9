package main

import "fmt"

func add(a, b []int32, p int) []int32 {
	if len(a) < len(b) {
		a, b = b, a
	}
	var storeDigit int32 = 0
	ans := make([]int32, len(a)+1)
	var c int32
	for i := 0; i < len(a); i++ {
		if i < len(b) {
			c = a[i] + b[i] + storeDigit
		} else {
			c = a[i] + storeDigit
		}
		ans[i] = c % int32(p)
		storeDigit = c / int32(p)
	}
	if storeDigit == 1 {
		ans[len(a)] = 1
	} else {
		ans = ans[:len(a)]
	}
	return ans
}

func main() {
	a := []int32{0, 2, 1, 2}
	b := []int32{1, 2, 2, 1}
	fmt.Println(add(a, b, 3))

	a = []int32{6, 2, 1, 1}
	b = []int32{5, 3, 0, 5}
	fmt.Println(add(a, b, 7))
}
