package main

import (
	"fmt"
	"math"
	"sort"
)

func Divisors(x int) []int {
    divisors := make(map[int]int)
    maxv := int(math.Sqrt(float64(x))) + 1
    for i := 1; i < maxv; i++ {
        if x % i == 0 {
            divisors[i] = 1
            divisors[x/i] = 1
        }
    }
    result := make([]int, 0, len(divisors))
    for idx := range divisors {
        result = append(result, idx)
    }
    sort.Ints(result)
    return result
}

func getEdges(divisors []int) [][2]int {
    result := make([][2]int, 0)
    for i := 0; i < len(divisors); i++ {
        for j := 0; j < i; j++ {
            if divisors[i] % divisors[j] == 0 {
                divisorBetween := false
                for k := j + 1; k < i; k++ {
                    if divisors[i] % divisors[k] == 0 && divisors[k] % divisors[j] == 0 {
                        divisorBetween = true
                        break
                    }
                }
                if !divisorBetween {
                    result = append(result, [2]int{divisors[i], divisors[j]})
                }
            }
        }
    }
    return result
}

func main() {
    var x int;
    fmt.Scan(&x);
    divisors := Divisors(x)
    edges := getEdges(divisors)

    fmt.Println("graph {")
    for i := len(divisors) - 1; i >= 0; i-- {
        fmt.Printf("    %d\n", divisors[i])
    }
    for i := len(edges) - 1; i >= 0; i-- {
        fmt.Printf("    %d--%d\n", edges[i][0], edges[i][1])
    }
    fmt.Println("}")
}