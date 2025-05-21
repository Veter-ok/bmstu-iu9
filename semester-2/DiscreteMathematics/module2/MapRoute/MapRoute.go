package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
    V []*Node
}

type Node struct {
    x, y, idx, index, m, dist int
    parent *Node
}

type PriorityQueue struct {
    heap []*Node;
    count, cap int
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
    for ;i > 0 && queue.heap[(i - 1)/2].dist > queue.heap[i].dist; i = (i - 1)/2 {
        queue.heap[(i - 1) / 2], queue.heap[i] = queue.heap[i], queue.heap[(i - 1) / 2]
        queue.heap[i].index = i
    }
    queue.heap[i].index = i
}

func (queue *PriorityQueue) DecreaseKey(i, dist int) {
    queue.heap[i].dist = dist
    for ; i > 0 && queue.heap[(i - 1)/2].dist > queue.heap[i].dist; i = (i - 1)/2 {
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
            if (l < queue.count) && queue.heap[i].dist > queue.heap[l].dist {
                i = l
            }
            if (r < queue.count) && queue.heap[i].dist > queue.heap[r].dist {
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

func Relax(u, v *Node, w int) bool {
    changed := u.dist + w < v.dist
    if changed  {
        v.dist = u.dist + w
        v.parent = u 
    }
    return changed
}

func Dijkstra(graph Graph, s *Node, n int) {
    queue := InitPriorityQueue(n*n)
    s.dist = s.m
    for _, v := range graph.V {
        queue.Insert(v)
    }
    for !queue.QueueEmpty() {
        v := queue.ExtractMin()
        v.index = -1
        if v.x < n - 1 && graph.V[v.idx+1].index != -1 && Relax(v, graph.V[v.idx+1], graph.V[v.idx+1].m) {
            queue.DecreaseKey(graph.V[v.idx+1].index, graph.V[v.idx+1].dist)
        }
        if v.x > 0 && graph.V[v.idx-1].index != - 1 && Relax(v, graph.V[v.idx-1], graph.V[v.idx-1].m) {
            queue.DecreaseKey(graph.V[v.idx-1].index, graph.V[v.idx-1].dist)
        }
        if v.y > 0 && graph.V[v.idx-n].index != -1 && Relax(v, graph.V[v.idx-n], graph.V[v.idx-n].m) {
            queue.DecreaseKey(graph.V[v.idx-n].index, graph.V[v.idx-n].dist)
        }
        if v.y < n - 1 && graph.V[v.idx+n].index != -1 && Relax(v, graph.V[v.idx+n], graph.V[v.idx+n].m) {
            queue.DecreaseKey(graph.V[v.idx+n].index, graph.V[v.idx+n].dist)
        }
    }
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    n, _ := strconv.Atoi(scanner.Text())

    V := make([]*Node, n*n)
    for i, y := 0, 0; y < n; y++ {
        scanner.Scan()
        parts := strings.Fields(scanner.Text())
        for x := 0; x < n; x++ {
            m, _ := strconv.Atoi(parts[x])
            V[i] = &Node{x, y, i, 0, m, 1_000_000, nil}
            i++
        }
    }

    graph := &Graph{V}
    Dijkstra(*graph, graph.V[0], n)
    fmt.Println(graph.V[n*n - 1].dist)
}