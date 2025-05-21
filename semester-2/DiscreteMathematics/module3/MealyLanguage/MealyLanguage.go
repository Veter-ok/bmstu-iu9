package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stateKey struct {
	q int
	s string
}

func searchAllPaths(f [][]string, delta [][]int, paths *map[string]bool,  
                    visited *map[stateKey]bool, curString string, q0, M int) {
    if len(curString) > M{
        return
    }
    key := stateKey{q0, curString}
    if (*visited)[key] {
        return
    }
    (*visited)[key] = true
    (*paths)[curString] = true
    for lIdx := 0; lIdx < 2; lIdx++ {
        signal := f[q0][lIdx]
        new_q := delta[q0][lIdx]
        if q0 == new_q && signal == "-" {
            continue
        }
        if signal != "-" {
            searchAllPaths(f, delta, paths, visited, curString+signal, new_q, M)
        }else{
            searchAllPaths(f, delta, paths, visited, curString, new_q, M)
        }
    }
}

func main(){
    reader := bufio.NewReader(os.Stdin)
    var n, M, q0 int

    line, _ := reader.ReadString('\n')
    n, _ = strconv.Atoi(strings.TrimSpace(line))
    reader.ReadString('\n')

    delta := make([][]int, n)
    for i := 0; i < n; i++ {
        delta[i] = make([]int, 2)
        line, _ = reader.ReadString('\n')
        parts := strings.Fields(line)
        for j := 0; j < 2; j++ {
            delta[i][j], _ = strconv.Atoi(parts[j])
        }
    }

    reader.ReadString('\n')
    f := make([][]string, n)
    for i := 0; i < n; i++ {
        line, _ = reader.ReadString('\n')
        f[i] = strings.Fields(line)
    }

    reader.ReadString('\n')
    line, _ = reader.ReadString('\n')
    q0, _ = strconv.Atoi(strings.TrimSpace(line))
    line, _ = reader.ReadString('\n')
    M, _ = strconv.Atoi(strings.TrimSpace(line))


    paths := make(map[string]bool)
    visited := make(map[stateKey]bool)
    searchAllPaths(f, delta, &paths, &visited, "", q0, M)
    for s := range paths {
        fmt.Printf("%s ", s)
    }
}