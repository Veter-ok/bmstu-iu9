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
    E []*Edge
    findByIdx map[int]*Node
    adj map[*Node][]*Node
}

type Node struct {
    idx,dist int
    parent *Node
}

type Edge struct {
    u, v *Node
}

func NewGraph(n int) *Graph {
    return &Graph{
        V: make([]*Node, 0),
        E: make([]*Edge, 0),
        findByIdx: map[int]*Node{},
        adj: make(map[*Node][]*Node),
    }
}

func Relax(u, v *Node, w int) bool {
    changed := u.dist + w > v.dist
    if changed  {
        v.dist = u.dist + w
        v.parent = u 
    }
    return changed
}

func (graph *Graph) BellmanFord(s *Node) {
    s.dist = 0
    for i := 1; i < len(graph.V); i++ {
        for _, edge := range graph.E {
            Relax(edge.u, edge.v, 1)
        }
    }
}


func (graph *Graph) HasCycle() bool {
    color := make(map[*Node]int)
    var dfs func(*Node) bool
    dfs = func(u *Node) bool {
        color[u] = 1
        for _, v := range graph.adj[u] {
            if color[v] == 1 {
                return true
            }
            if color[v] == 0 && dfs(v) {
                return true
            }
        }
        color[u] = 2 
        return false
    }
    for _, node := range graph.V {
        if color[node] == 0 {
            if dfs(node) {
                return true
            }
        }
    }
    return false
}

func main() {
    reader := bufio.NewReader(os.Stdin)

    haveParent := make(map[*Node]bool)
    var n int
    fmt.Fscanf(reader, "%d\n", &n)
    graph := NewGraph(n)
    r := 1
    for i := 0; i < n; i++ {
        line, _ := reader.ReadString('\n')
        parts := strings.Fields(line)
        k, _ := strconv.Atoi(parts[0])
        if _, ok := graph.findByIdx[r]; !ok {
            node := &Node{r, -1000000, nil}
            graph.V = append(graph.V, node)
            graph.findByIdx[node.idx] = node
        }
        node := graph.findByIdx[r]
        r++
        if k != 0 {
           for j := 0; j < k; j++ {
                c, _ := strconv.Atoi(parts[j+1])
                if _, ok := graph.findByIdx[c]; !ok {
                    node2 := &Node{c, -1000000, nil}
                    graph.V = append(graph.V, node2)
                    graph.findByIdx[c] = node2
                }
                edge := &Edge{graph.findByIdx[c], node}
                graph.E = append(graph.E, edge)
                graph.adj[graph.findByIdx[c]] = append(graph.adj[graph.findByIdx[c]], node)
                haveParent[node] = true
            } 
        }
    }

    if graph.HasCycle() {
        fmt.Println("cycle")
        return
    }

    startNodes := make([]*Node, 0)
    for _, v := range graph.V { 
        if ok, _ := haveParent[v]; !ok {
            startNodes = append(startNodes, v)  
        }
    }

    for _, v := range startNodes {
        graph.BellmanFord(v)
    }

    max := 0
    for _, v := range graph.V { 
        if v.dist > max && v.dist != -1000000 {
            max = v.dist
        }
    }
    fmt.Println(max + 1)
}
