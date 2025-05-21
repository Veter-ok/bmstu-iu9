package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
    adj [][]int
    mark []int
}

func NewGraph(n int) *Graph {
    return &Graph{
        adj: make([][]int, n),
        mark: make([]int, n),
    }
}

func (g *Graph) AddEdge(u, v int) {
    g.adj[u] = append(g.adj[u], v)
    g.adj[v] = append(g.adj[v], u)
}

func dfs(graph *Graph, v, color int, countEdges *int, component *[]int) {
    graph.mark[v] = color
    *component = append(*component, v)
    for _, u := range graph.adj[v] {
        if graph.mark[u] == 0 {
            (*countEdges)++
            dfs(graph, u, color, countEdges, component)
        }
    } 
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    var n, m int
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

    var maxCompLen, maxCountEdges, colorBestComp int
    for i := 0; i < n; i++ {
        if graph.mark[i] == 0 {
            var component []int
            var countEdges int
            dfs(graph, i, i+1, &countEdges, &component)
            if maxCompLen < len(component) ||
               maxCompLen == len(component) && maxCountEdges < countEdges ||
               maxCompLen == len(component) && maxCountEdges == countEdges && colorBestComp > component[0] {
                maxCompLen = len(component)
                colorBestComp = i+1
                maxCountEdges = countEdges
            }
        }
    }
    fmt.Println("graph {")
    for i := 0; i < n; i++ {
        if colorBestComp == graph.mark[i] {
            fmt.Printf("    %d [color=red]\n", i)
        }else {
            fmt.Printf("    %d\n", i)
        }
        for _, elem := range graph.adj[i]{
            if elem > i {
                if colorBestComp == graph.mark[i] {
                    fmt.Printf("    %d--%d [color=red]\n", i, elem)
                }else{
                    fmt.Printf("    %d--%d\n", i,elem )
                }
            }
        }
    }
    fmt.Println("}")
}