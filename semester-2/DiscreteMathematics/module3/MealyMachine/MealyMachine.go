package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dfs(q0, m, n int, idx *int, states_in_new_idx, new_idx []int, delta [][]int){
    states_in_new_idx[*idx] = q0
    new_idx[q0] = *idx
    (*idx)++
    for i := 0; i < m; i++{
        if new_idx[delta[q0][i]] == -1 {
            dfs(delta[q0][i], m, n, idx, states_in_new_idx, new_idx, delta)
        }
    }
}

func main(){
    reader := bufio.NewReader(os.Stdin)
    var n, m, q0 int
    fmt.Scanf("%d\n%d\n%d", &n, &m, &q0)

    delta := make([][]int, n)
    for i := 0; i < n; i++ {
        delta[i] = make([]int, m)
        line, _ := reader.ReadString('\n')
        parts := strings.Fields(line)
        for j := 0; j < m; j++{
            delta[i][j], _ = strconv.Atoi(parts[j])
        }
    }

    f := make([][]string, n)
    for i := 0; i < n; i++ {
        f[i] = make([]string, m)
        line, _ := reader.ReadString('\n')
        parts := strings.Fields(line)
        for j := 0; j < m; j++ {
            f[i][j] = parts[j]
        }
    }

    idx := 0
    new_idx := make([]int, n)
    states_in_new_idx := make([]int, n)
    for i := 0; i < n; i++ {
        new_idx[i] = -1
    }
    dfs(q0, m, n, &idx, states_in_new_idx, new_idx, delta)

    fmt.Printf("%d\n%d\n%d\n", n, m, 0)
    for i := 0; i < n; i++ {
        old_state := states_in_new_idx[i]
        for j := 0; j < m; j++ {
            fmt.Printf("%d ", new_idx[delta[old_state][j]])
        }
        fmt.Println()
    }
    for i := 0; i < n; i++ {
        old_state := states_in_new_idx[i]
        for j := 0; j < m; j++ {
            fmt.Printf("%s ", f[old_state][j])
        }
        fmt.Println()
    }
}