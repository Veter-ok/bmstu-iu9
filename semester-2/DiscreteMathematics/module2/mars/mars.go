package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Graph struct {
	adj [][]int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func NewGraph(n int) *Graph {
	return &Graph{
		adj: make([][]int, n),
	}
}

func (g *Graph) AddEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
}

func isLexSmaller(group1, group2 []int) bool {
	for i := 0; i < len(group1) && i < len(group2); i++ {
		if group1[i] < group2[i] {
			return true
		} else if group1[i] > group2[i] {
			return false
		}
	}
	return len(group1) < len(group2)
}

func findAllColorings(graph *Graph, v int, color []int, allColorings *[][]int) {
	if v == len(graph.adj) {
		coloring := make([]int, len(color))
		copy(coloring, color)
		*allColorings = append(*allColorings, coloring)
		return
	}
	color[v] = 1
	valid := true
	for _, u := range graph.adj[v] {
		if color[u] == color[v] {
			valid = false
			break
		}
	}
	if valid {
		findAllColorings(graph, v+1, color, allColorings)
	}
	color[v] = -1
	valid = true
	for _, u := range graph.adj[v] {
		if color[u] == color[v] {
			valid = false
			break
		}
	}
	if valid {
		findAllColorings(graph, v+1, color, allColorings)
	}
	color[v] = 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	graph := NewGraph(n)

	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		fields := strings.Fields(line)
		for j := 0; j < n; j++ {
			if j > i && fields[j] == "+" {
				graph.AddEdge(i, j)
			}
		}
	}
	color := make([]int, n)
	allColorings := [][]int{}
	findAllColorings(graph, 0, color, &allColorings)

	if len(allColorings) == 0 {
		fmt.Println("No solution")
		return
	}
	bestGroup := []int{}
	minDiff := n + 1

	for _, coloring := range allColorings {
		var group1, group2 []int
		for i := 0; i < n; i++ {
			if coloring[i] == 1 {
				group1 = append(group1,i)
			} else {
				group2 = append(group2,i)
			}
		}
		if len(group1) > len(group2) {
			group1, group2 = group2, group1
		}
		diff := abs(len(group1) - len(group2))
		if diff < minDiff || (diff == minDiff && isLexSmaller(group1, bestGroup)) {
			minDiff = diff
			bestGroup = group1
		}
	}
	for _, v := range bestGroup {
		fmt.Print(v+1, " ")
	}
}