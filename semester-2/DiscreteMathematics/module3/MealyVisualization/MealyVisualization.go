package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
    reader := bufio.NewReader(os.Stdin)
    var n, m, q0 int
    alphabet := "abcdefghijklmnopqrstuvwxyz"
    fmt.Scanf("%d\n%d\n%d", &n, &m, &q0)

    delta := make([][]int, n)
    for i := 0; i < n; i++ {
        delta[i] = make([]int, m)
        line, _ := reader.ReadString('\n')
        parts := strings.Fields(line)
        for j := 0; j < m; j++ {
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

    fmt.Println("digraph {\n    rankdir = LR")
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            fmt.Printf("    %d -> %d [label = \"%s(%s)\"]\n", i, delta[i][j], string(alphabet[j]), f[i][j])
        }
    }
    fmt.Println("}")
}