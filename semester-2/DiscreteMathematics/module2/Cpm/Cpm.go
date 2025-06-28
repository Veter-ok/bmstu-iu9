package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Queue struct {
    data []*Node
    cap, count, head, tail int
}

type Graph struct {
    V, mainV []*Node
    VbyName map[string]*Node
    adj map[*Node][]*Node
}

type Node struct {
    time0, time1, color, mark int
    parents []*Node
    name string
}

func InitQueue(cap int) *Queue {
    return &Queue{make([]*Node, cap), cap, 0, 0, 0}
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

func (graph *Graph) dfs(node *Node) []*Node {
    node.mark = 1
    ans := make([]*Node, 0)
    for _, v := range graph.adj[node] {
        if v.mark == 0 {
            u := graph.dfs(v)
            if len(u) > 0 {
                ans = append(ans, u...)
            }
        }
        if v.mark == 1 {
            ans = append(ans, v)
        }
    }
    node.mark = 2
    return ans
}

func (graph *Graph) bfs(n int) {
    queue := InitQueue(n)
    for _, w := range graph.mainV {
        if w.time1 == -1 && w.color != 2 {
            w.time1 = w.time0
            queue.Enqueue(w)
        }
        for queue.count > 0 {
            v := queue.Dequeue()
            for _, next := range graph.adj[v] {
                if next.color == 2 {
                    continue
                }
                new_time := v.time1 + next.time0
                if next.time1 == -1 || next.time1 < new_time {
                    next.time1 = new_time
                    next.parents = []*Node{v}
                }else if next.time1 == new_time {
                    next.parents = append(next.parents, v)
                }
                queue.Enqueue(next)
            }
        }
    }
}

func (graph *Graph) colorGraph(node *Node, color int) {
    node.color = color
    var nodes []*Node
    if color == 1 {
        nodes = node.parents
    }else{
        nodes = graph.adj[node]
    }
    for _, v := range nodes {
        if v.color != color {
            graph.colorGraph(v, color)
        }
    }
}

func getNamesAndNums(str string) ([][]int, [][]string) {
    names, nums := make([][]string, 1), make([][]int, 1)
    name, num, n, nameOver := "", 0, 0, false
    for i := 0; i < len(str); i++ {
        if str[i] == '(' || str[i] == ')'{
            nameOver = true
            continue
        }else if str[i] == '<' ||  str[i] == ';' {
            names[n] = append(names[n], name)
            nums[n] = append(nums[n], num / 10)
            name, num, nameOver = "", 0, false
            if str[i] == ';' {
                names = append(names, make([]string, 0))
                nums = append(nums, make([]int, 0))
                n++
            }
            continue
        }
        if nameOver {
            num = (num + (int(str[i]) - 48)) * 10
        }else {
            name += string(str[i])
        }
    }
    if name != "" {
        names[n] = append(names[n], name)
        nums[n] = append(nums[n], num / 10)
    }
    return nums, names
}

func (graph *Graph) addNodesAndEdges(node, root *Node) *Node {
    if graph.VbyName[node.name] == nil {
        graph.V = append(graph.V, node)
        graph.VbyName[node.name] = node
    }else{
        node = graph.VbyName[node.name]
    }
    if root != nil {
        graph.adj[root] = append(graph.adj[root], node)
    }
    return node
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    graph := &Graph{make([]*Node, 0),make([]*Node, 0), make(map[string]*Node), make(map[*Node][]*Node)}
    sentence := ""
    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            break
        }
        sentence += line
    }

    sentence = strings.ReplaceAll(sentence, " ", "")
    nums, names := getNamesAndNums(sentence)
    for i := 0; i < len(nums); i++ {
        var root *Node
        for j := 0; j < len(nums[i]); j++ {
            node := &Node{nums[i][j], -1, 0, 0, nil, names[i][j]}
            root = graph.addNodesAndEdges(node, root)
            if j == 0 {
                graph.mainV = append(graph.mainV, root)
            }
        }
    }

    for _, v := range graph.mainV {
        if v.mark != 2 {
            nodes := graph.dfs(v)
            for _, u := range nodes {
                graph.colorGraph(u, 2)
            }
        }
    }
    graph.bfs(len(graph.V))
    maxNodes := make([]*Node, 0)
    for _, v := range graph.V {
        if v.color != 2 {
            if len(maxNodes) == 0 || maxNodes[0].time1 < v.time1 {
                maxNodes = []*Node{v}
            }
            if maxNodes[0].time1 == v.time1 {
                maxNodes = append(maxNodes, v)
            }
        }
    }
    for _, v := range maxNodes {
        graph.colorGraph(v, 1)
    } 

    fmt.Println("digraph {")
    for _, v := range graph.V {
        if v.color == 1 {
            fmt.Printf("    %s [label = \"%s(%d)\", color = \"red\"]\n", v.name, v.name, v.time0)
        }else if v.color == 2 {
            fmt.Printf("    %s [label = \"%s(%d)\", color = \"blue\"]\n", v.name, v.name, v.time0)
        }else {
            fmt.Printf("    %s [label = \"%s(%d)\"]\n", v.name, v.name, v.time0)
        }
    }
    for _, v := range graph.V { 
        for _, u := range graph.adj[v] {
            if v.color == 1 && u.color == 1 && v.time1 + u.time0 == u.time1 {
                fmt.Printf("    %s -> %s [color = \"red\"]\n", v.name, u.name)
            }else if v.color == 2 && u.color == 2 {
                fmt.Printf("    %s -> %s [color = \"blue\"]\n", v.name, u.name)
            }else {
                fmt.Printf("    %s -> %s\n", v.name, u.name)
            }
        }
    }
    fmt.Println("}")
}