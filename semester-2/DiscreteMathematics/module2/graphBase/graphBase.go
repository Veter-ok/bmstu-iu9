package main

import "fmt"

var time, count = 1, 1

type Stack struct {
    nodes []*Node
    top int
}

type Node struct {
    idx, comp, low, t1 int
}

type Graph struct {
    adj [][]*Node
    V []*Node
}

func NewGraph(n int) *Graph {
    return &Graph{
        V: make([]*Node, n),
        adj: make([][]*Node, n),
    }
}

func InitStack(cap int) *Stack {
    return &Stack{
        nodes: make([]*Node, cap),
        top: 0,
    }
}

func (stack *Stack) Push(x *Node) {
    stack.nodes[stack.top] = x
    stack.top++
}

func (stack *Stack) Pop() *Node {
    stack.top--
    x := stack.nodes[stack.top]
    stack.nodes[stack.top] = nil
    return x
}

func Tarjan(graph *Graph){
    stack := InitStack(len(graph.V))
    for _, v := range graph.V {
        if v.t1 == 0 {
            VisitVertex_Tarjan(graph, v, stack)
        }
    }
}

func VisitVertex_Tarjan(graph *Graph, v *Node, stack *Stack) {
    v.t1, v.low = time, time
    time++
    stack.Push(v)
    for _, u := range graph.adj[v.idx] {
        if u.t1 == 0 {
            VisitVertex_Tarjan(graph, u, stack)
        } 
        if u.comp == 0 && v.low > u.low {
            v.low = u.low
        }
    }
    if v.t1 == v.low {
        var u *Node
        for {
            u = stack.Pop()
            u.comp = count
            if u.idx == v.idx {
                break
            }
        }
        count++
    }
}

func main() {
    var n, m int
    fmt.Scanf("%d\n%d", &n, &m)

    graph := NewGraph(n)
    for i := 0; i < n; i++ {
        graph.V[i] = &Node{i, 0, 0, 0}
    }
    for i := 0; i < m; i++ {
        var u, v int
        fmt.Scanf("%d %d", &u, &v)
        graph.adj[u] = append(graph.adj[u], graph.V[v])
    }
    Tarjan(graph)
    min := make([]int, count)
    used := make([]bool, count)
    for i := 0; i < count; i++ {
        used[i] = true
        min[i] = 100000
    }
    for _, v := range graph.V {
        if min[v.comp] > v.idx {
            min[v.comp] = v.idx
        }
    }
    for _, u := range graph.V {
        for _, v := range graph.adj[u.idx] {
            if u.comp != v.comp {
                used[v.comp] = false
            }
        }
    }
    for i := 1; i < count; i++ {
        if used[i] {
            fmt.Println(min[i])
        }
    }
}