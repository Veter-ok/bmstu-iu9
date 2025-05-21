package main

import (
	"fmt"
	"sort"
)

type Queue struct {
    data []*Node
    cap, count, head, tail int
}

type Node struct {
    idx int
    colors []int
}

type Graph struct {
    adj [][][2]int
}

func NewGraph(n int) *Graph {
    return &Graph{
        adj: make([][][2]int, n),
    }
}

func (g *Graph) AddEdge(u, v, c int) {
    g.adj[u] = append(g.adj[u], [2]int{v, c})
    g.adj[v] = append(g.adj[v], [2]int{u, c})
}

func InitQueue(cap int) *Queue {
    return &Queue{
        data: make([]*Node, cap),
        cap: cap,
        count: 0,
        head: 0,
        tail: 0,
    }
}

func (queue *Queue) Enqueue(x *Node) {
    queue.data[queue.tail] = x
    queue.tail ++
    if (queue.tail == queue.cap) {
        queue.tail = 0
    }
    queue.count++
}

func (queue *Queue) Dequeue() *Node {
    x := queue.data[queue.head]
    queue.head++
    if (queue.head == queue.cap) {
        queue.head = 0
    }
    queue.count--
    return x
}

func isLexSmaller(a, b []int) bool {
    for i := 0; i < len(a); i++ {
        if a[i] > b[i] {
            return false
        }
    }
    return true
}

func (graph *Graph) bfs(n int) (int, []int){
    dist := make([]int, n)
    paths := make([][]int, n)
    for i := range dist {
        dist[i] = -1
        paths[i] = []int{}
    }
    dist[0] = 0
    queue := InitQueue(n*10)
    queue.Enqueue(&Node{idx: 0, colors: []int{}})
    for queue.count > 0 {
        v := queue.Dequeue()
        for _, u := range graph.adj[v.idx] {
            next, color := u[0], u[1]
            newColors := append([]int{}, v.colors...)
            newColors = append(newColors, color)
            if dist[next] == -1 {
                dist[next] = dist[v.idx] + 1
                paths[next] = newColors
                queue.Enqueue(&Node{idx: next, colors: newColors})
            } else if dist[next] == dist[v.idx] + 1{
                if isLexSmaller(newColors, paths[next]) {
                    paths[next] = newColors
                    queue.Enqueue(&Node{idx: next, colors: newColors})
                }
            }
        }
    }
    return dist[n-1], paths[n-1]
}

func main(){
    var n, m int
    fmt.Scanf("%d %d", &n, &m)
    graph := NewGraph(n)
    for i := 0; i < m; i++ {
        var a, b, c int
        fmt.Scanf("%d %d %d", &a, &b, &c)
        graph.AddEdge(a - 1, b - 1, c)
    }
    for i := 0; i < n; i++ {
        sort.Slice(graph.adj[i], func(j, k int) bool {
            return graph.adj[i][j][1] < graph.adj[i][k][1]
        })
    }
    dist, paths := graph.bfs(n)
    fmt.Println(dist)
    for _, i := range paths {
        fmt.Printf("%d ", i)
    }
}