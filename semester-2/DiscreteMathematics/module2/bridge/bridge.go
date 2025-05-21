package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
    adj []*Node
    h []int
    d []int
    mark []int
}

type Node struct {
    idx int
    next *Node
}

func NewGraph(n int) *Graph {
    return &Graph{
        adj: make([]*Node, n),
        h: make([]int, n),
        d: make([]int, n),
        mark: make([]int, n),
    }
}

func (g *Graph) AddEdge(u, v int) {
    g.adj[u] = &Node{v, g.adj[u]}
	g.adj[v] = &Node{u, g.adj[v]}
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func dfs(graph *Graph, v, p int, bridge *int){
    graph.mark[v] = 1
    if (p == -1) {
        graph.d[v], graph.h[v] = 0, 0
    }else{
        graph.d[v], graph.h[v] = graph.h[p] + 1, graph.h[p] + 1
    }
    for u := graph.adj[v]; u != nil; u = u.next {
        if u.idx != p {
            if graph.mark[u.idx] == 1 {
                graph.d[v] = min(graph.d[v], graph.h[u.idx])
            }else{
                dfs(graph, u.idx, v, bridge)
                graph.d[v] = min(graph.d[v], graph.d[u.idx])
                if graph.h[v] < graph.d[u.idx] {
                    (*bridge)++
                }
            }
        }
    }
}

func main(){
    reader := bufio.NewReader(os.Stdin)
    var n, m, ans int
    fmt.Fscanf(reader, "%d\n", &n)
    fmt.Fscanf(reader, "%d\n", &m)
    graph := NewGraph(n)

    for i := 0; i < m; i++ {
        line, _ := reader.ReadString('\n')
        line = strings.TrimSpace(line)
        nums := strings.Fields(line);
        num1, _ := strconv.Atoi(nums[0])
        num2, _ := strconv.Atoi(nums[1])
        graph.AddEdge(num1, num2)
    }
    for i := 0; i < n; i++ {
		if graph.mark[i] == 0 {
			dfs(graph, i, -1, &ans)
		}
	}
    fmt.Println(ans);
}