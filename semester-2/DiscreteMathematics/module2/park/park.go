package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Graph struct {
    V []*Node
    E []*Edge
}

type Node struct {
    idx, x, y int
}

type Edge struct {
    u, v *Node
    w float64
}

type Set struct {
    parent []int
    rank   []int
}

func NewGraph(n int) *Graph {
    return &Graph{
        V: make([]*Node, n),
        E: []*Edge{},
    }
}

func (g *Graph) AddEdge(u, v int, w float64) {
    g.E = append(g.E, &Edge{g.V[u], g.V[v], w})
}

func NewSet(size int) *Set {
    set := &Set{
        parent: make([]int, size),
        rank:   make([]int, size),
    }
    for i := range set.parent {
        set.parent[i] = i
    }
    return set
}

func (set *Set) find(v int) int {
    if set.parent[v] != v {
        set.parent[v] = set.find(set.parent[v])
    }
    return set.parent[v]
}

func (set *Set) union(u, v int) {
    uRoot := set.find(u)
    vRoot := set.find(v)
    if uRoot == vRoot {
        return
    }
    if set.rank[uRoot] < set.rank[vRoot] {
        set.parent[uRoot] = vRoot
    } else if set.rank[uRoot] > set.rank[vRoot] {
        set.parent[vRoot] = uRoot
    } else {
        set.parent[vRoot] = uRoot
        set.rank[uRoot]++
    }
}

func MST_Kruskal(graph *Graph) float64 {
    sort.Slice(graph.E, func(i, j int) bool {
        return graph.E[i].w < graph.E[j].w
    })
    totalWeight := float64(0)
    var E []Edge
    set := NewSet(len(graph.V))

    for i := 0; i < len(graph.E) - 1 && len(E) < len(graph.V) - 1; i++ {
        u, v := graph.E[i].u, graph.E[i].v
        if set.find(u.idx) != set.find(v.idx) {
            E = append(E, *graph.E[i]) 
            set.union(u.idx, v.idx)
            totalWeight += graph.E[i].w
        }
    }
    return totalWeight
}

func main(){
    reader := bufio.NewReader(os.Stdin)
    var n int
    fmt.Fscanf(reader, "%d\n", &n)
    graph := NewGraph(n)

    for i := 0; i < n; i++ {
        line, _ := reader.ReadString('\n')
        line = strings.TrimSpace(line)
        nums := strings.Fields(line);
        num1, _ := strconv.Atoi(nums[0])
        num2, _ := strconv.Atoi(nums[1])
        graph.V[i] = &Node{i, num1, num2}
    }
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if i != j {
                x := float64(graph.V[i].x - graph.V[j].x)
                y := float64(graph.V[i].y - graph.V[j].y)
                w := float64(math.Sqrt(x*x + y*y))
                graph.AddEdge(i, j, w)
            }
        }
    }
    totalWeight := MST_Kruskal(graph)
    fmt.Printf("%.2f", totalWeight)
}