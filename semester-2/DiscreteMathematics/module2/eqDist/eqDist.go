package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Graph struct {
    V []*Node
    
    adj [][]*Node
}

type Node struct {
    idx, w int
    str string
}

func getNameAndNum(str string) (int, string) {
    name, num, nameOver := "", 0, false
    str = strings.TrimSpace(str)
    for i := 0; i < len(str); i++ {
        if str[i] == '(' {
            nameOver = true
            continue
        }else if str[i] == ')' {
            break
        }
        if nameOver {
            num = (num + (int(str[i]) - 48)) * 10
        }else{
            name += string(str[i])
        }
    }
    return num / 10, name
}

func main() {
    reader := bufio.NewReader(os.Stdin)

    graph := &Graph{make([]*Node, 0), make([][]*Node, 0)}
    sentence, _ := reader.ReadString(';')
    nodes := strings.Split(sentence, "<")
    for i := 0; i < len(nodes); i++ {
        num, name := getNameAndNum(nodes[i])
        graph.V = append(graph.V, &Node{i, num, name})
    }

    for _, n := range graph.V {
        fmt.Println(n.str, n.w)
    }
}