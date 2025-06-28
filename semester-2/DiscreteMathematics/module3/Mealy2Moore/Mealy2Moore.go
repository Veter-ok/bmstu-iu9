package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MooreState struct {
    idx int
    outputSymbol string
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    var k1, k2, n int

    fmt.Fscanf(reader, "%d\n", &k1)
    X := make([]string, 0)
    line, _ := reader.ReadString('\n')
    x := strings.Fields(line)
    for _, val := range x {
        X = append(X, val)
    }

    fmt.Fscanf(reader, "%d\n", &k2)
    Y := make([]string, 0)
    line, _ = reader.ReadString('\n')
    y := strings.Fields(line)
    for _, val := range y {
        Y = append(Y, val)
    }

    fmt.Fscanf(reader, "%d\n", &n)
    delta := make(map[int]map[string]int)
    for i := 0; i < n; i++ {
        line, _ = reader.ReadString('\n')
        parts := strings.Fields(line)
        delta[i] = make(map[string]int)
        for j := 0; j < k1; j++ {
            delta[i][X[j]], _ = strconv.Atoi(parts[j])
        }
    }

    phi := make(map[int]map[string]string)
    for i := 0; i < n; i++ {
        line, _ = reader.ReadString('\n')
        parts := strings.Fields(line)
        phi[i] = make(map[string]string)
        for j := 0; j < k1; j++ {
            sn, _ := strconv.Atoi(parts[j])
            phi[i][X[j]] = Y[sn]
        }
    }

    pairs := make(map[MooreState]int)
    new_Q := make([]MooreState, 0)
    for i := 0; i < n; i++ {
        for _, x := range X {
            mooreState := MooreState{delta[i][x], phi[i][x]}
            if _, ok := pairs[mooreState]; !ok {
                new_Q = append(new_Q, mooreState)
                pairs[mooreState] = len(pairs)
            }
        }
    }   

    fmt.Println("digraph {\n    rankdir = LR")
    for idx, state := range new_Q {
        fmt.Printf("    %d [label=\"(%d,%s)\"]\n", idx, state.idx, state.outputSymbol)
    }
    for idx, state := range new_Q {
        for i := 0; i < k1; i++ {
            mooreState := MooreState{delta[state.idx][X[i]], phi[state.idx][X[i]]}
            fmt.Printf("    %d -> %d [label=\"%s\"]\n", idx, pairs[mooreState], X[i])
        }
    }
    fmt.Println("}")
}