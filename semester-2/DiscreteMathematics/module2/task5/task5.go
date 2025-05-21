package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
    adj []*Node
}

type Node struct {
    idx int
    next *Node
}

type Pair struct {
    dist, node int
}

func NewGraph(n int) *Graph {
    return &Graph{
        adj: make([]*Node, n),
    }
}

func (g *Graph) AddEdge(u, v int) {
    g.adj[u] = &Node{v, g.adj[u]}
    g.adj[v] = &Node{u, g.adj[v]}
}

func del(q *[]Pair, p Pair) {
    for i, v := range *q {
        if v == p {
            *q = append((*q)[:i], (*q)[i+1:]...)
        }
    }
}

func dijkstra(graph *Graph, s, n int) []int {
    dist := make([]int, n)
    for i := range dist {
        dist[i] = math.MaxInt32
    }
    dist[s] = 0
    q := []Pair{{0, s}}
    for len(q) != 0 {
        v := q[0].node
        q = q[1:]
        for node := graph.adj[v]; node != nil; node = node.next {
            u := node.idx
            if dist[u] > dist[v] + 1 {
                del(&q, Pair{dist[u], u})
                dist[u] = dist[v] + 1
                q = append(q, Pair{dist[u], u})
            }
        }
    }
    return dist
}

func main(){
    reader := bufio.NewReader(os.Stdin)
    var n, m, k int
    var v []int
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
    fmt.Fscanf(reader, "%d\n", &k)
    for i := 0; i < k; i++ {
        var num int
        fmt.Fscanf(reader, "%d ", &num)
        v = append(v, num)
    }

    perfect_dist := make([]int, n)
    for i := 0; i < n; i++ {
        perfect_dist[i] = -1
    }

    for i := 0; i < k; i++ {
        dist := dijkstra(graph, v[i], n)
        for j := 0; j < n; j++ {
            if perfect_dist[j] == -1 && dist[j] != math.MaxInt32 {
                perfect_dist[j] = dist[j]
            } else if perfect_dist[j] != dist[j] {
                perfect_dist[j] = 0
            }
        }
    }
    noAns := true
    for i := 0; i < n; i++ {
        if perfect_dist[i] > 0 {
            noAns = false
            fmt.Printf("%d ", i)
        }
    }
    if noAns {
        fmt.Println("-")
    }
}