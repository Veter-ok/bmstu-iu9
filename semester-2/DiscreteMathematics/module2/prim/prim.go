package main

import (
	"fmt"
	"math"
)

type Graph struct {
    V []*Node
    E []*Edge
}

type Edge struct {
    u, v *Node
    w int
}

type Node struct {
    index, key int
    value *Node
}

type PriorityQueue struct {
    heap []*Node;
    count, cap int
}

func NewGraph(n int) *Graph {
    nodes := make([]*Node, n)
    for i := range nodes {
        nodes[i] = &Node{-1, math.MaxInt32, nil}
    }
    return &Graph{
        V: nodes,
        E: []*Edge{},
    }
}

func (g *Graph) AddEdge(u, v int, w int) {
    g.E = append(g.E, &Edge{g.V[u], g.V[v], w})
    g.E = append(g.E, &Edge{g.V[v], g.V[u], w})
}

func InitPriorityQueue(n int) *PriorityQueue {
    return &PriorityQueue{
        heap: make([]*Node, n),
        count: 0,
        cap: n,
    }
}

func (queue *PriorityQueue) QueueEmpty() bool {
    return queue.count == 0
}

func (queue *PriorityQueue) Insert(x *Node) {
    i := queue.count
    queue.count++
    x.index = i
    queue.heap[i] = x;
    for ;i > 0 && queue.heap[(i - 1)/2].key < queue.heap[i].key; i = (i - 1)/2 {
        queue.heap[(i - 1) / 2], queue.heap[i] = queue.heap[i], queue.heap[(i - 1) / 2]
        queue.heap[i].index = i
    }
    queue.heap[i].index = i
}

func (queue *PriorityQueue) DecreaseKey(i, key int) {
    queue.heap[i].key = key
    for ; i > 0 && queue.heap[(i - 1)/2].key > queue.heap[i].key; i = (i - 1)/2 {
        queue.heap[(i - 1) / 2], queue.heap[i] = queue.heap[i], queue.heap[(i - 1) / 2]
        queue.heap[i].index = i
    }
    queue.heap[i].index = i
}

func (queue *PriorityQueue) ExtractMin() *Node {
    ptr := queue.heap[0]
    queue.count--
    ptr.index = -1
    if queue.count > 0 {
        queue.heap[0] = queue.heap[queue.count]
        queue.heap[0].index = 0
        queue.heap = queue.heap[:queue.count]
        i := 0
        for {
            l := 2*i + 1
            r := l + 1
            j := i
            if (l < queue.count) && queue.heap[i].key > queue.heap[l].key {
                i = l
            }
            if (r < queue.count) && queue.heap[i].key > queue.heap[r].key {
                i = r
            }
            if (i == j){
                break
            }
            queue.heap[i], queue.heap[j] = queue.heap[j], queue.heap[i]
            queue.heap[i].index = i
            queue.heap[j].index = j
        }
    }
    return ptr
}

func MST_Prim(graph *Graph) []*Edge {
    E := make([]*Edge, 0)
    for _, v := range graph.V {
        v.index = -1
    }
    queue := InitPriorityQueue(len(graph.V))
    for _, v := range graph.V {
        queue.Insert(v)
    }
    v := queue.ExtractMin()
    for !queue.QueueEmpty() {
        v.index = -2
        for _, e := range graph.E {
            var u *Node
            var a int
            if e.v == v {
                u, a = e.u, e.w
            }else if e.u == v {
                u, a = e.v, e.w
            }else{
                continue
            }
            if u.index == -1 {
                u.value, u.key = v, a
                queue.Insert(u)
            }else if u.index != -2 && a < u.key {
                u.value = v
                queue.DecreaseKey(u.index, a)
            } 
        }
        v = queue.ExtractMin()
        E = append(E, &Edge{v, v.value, v.key});
    }
    return E
}

func main() {
    var n, m int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &m)

    graph := NewGraph(n)
    for i := 0; i < m; i++ {
        var u, v, l int
        fmt.Scanf("%d %d %d", &u, &v, &l)
        graph.AddEdge(u, v, l)
    }
    E := MST_Prim(graph)
    totalWeight := 0
    for _, e := range E {
        totalWeight += e.w
    }
    fmt.Println(totalWeight)
}