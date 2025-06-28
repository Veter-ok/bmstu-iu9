package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type Node struct {
    formula string
    vars, depVars []string
    mark int
}

type Graph struct {
    V []*Node
    VbyName map[string]*Node
    E map[*Node][]*Node
}

type Tag int

type Lexem struct {
    Tag
    Image string
}

const (
    ERROR Tag = 1 << iota
    NUMBER
    VAR
    PLUS
    MINUS
    MUL
    DIV
    LPAREN
    RPAREN
    COMMA
    EQUAL
    EOF
)

func lexer(expr string) []Lexem {
    lexems := make([]Lexem, 0)
    for i := 0; i < len(expr); i++ {
        switch {
        case unicode.IsSpace(rune(expr[i])):
            continue;
        case expr[i] == '+':
            lexems = append(lexems, Lexem{PLUS, string(expr[i])})
        case expr[i] == '-':
            lexems = append(lexems, Lexem{MINUS, string(expr[i])})
        case expr[i] == '*':
            lexems = append(lexems, Lexem{MUL, string(expr[i])})
        case expr[i] == '/':
            lexems = append(lexems, Lexem{DIV, string(expr[i])})
        case expr[i] == '(':
            lexems = append(lexems, Lexem{LPAREN, string(expr[i])})
        case expr[i] == ')':
            lexems = append(lexems, Lexem{RPAREN, string(expr[i])})
        case expr[i] == '=':
            lexems = append(lexems, Lexem{EQUAL, string(expr[i])})
        case expr[i] == ',':
            lexems = append(lexems, Lexem{COMMA, string(expr[i])})
        case unicode.IsDigit(rune(expr[i])):
            j := i 
            for j < len(expr) && unicode.IsDigit(rune(expr[j])) {
                j++
            }
            lexems = append(lexems, Lexem{NUMBER, expr[i:j]})
            i = j - 1
        case unicode.IsLetter(rune(expr[i])):
            j := i 
            for j < len(expr) && (unicode.IsDigit(rune(expr[j])) || unicode.IsLetter(rune(expr[j]))) {
                j++
            }
            lexems = append(lexems, Lexem{VAR, expr[i:j]})
            i = j - 1;
        default:
            lexems = append(lexems, Lexem{ERROR, string(expr[i])})
        }
    }
    lexems = append(lexems, Lexem{EOF, "eof"})
    return lexems
}

func parser(formula string) ([]string, []string, bool) {
    lexems := lexer(formula)
    equal := false
    bracket, comma := 0, 0
    vars, depVars := make([]string, 0), make([]string, 0)
    operations := PLUS | MINUS | MUL | DIV
    for i := 0; i < len(lexems) - 1; i++ {
        curLex, nextLex := lexems[i], lexems[i+1]
        if curLex.Tag == EQUAL {
            if equal {
                return nil, nil, false
            }
            equal = true
        }
        if curLex.Tag == LPAREN {
            bracket++
        }else if curLex.Tag == RPAREN {
            bracket--
        }
        if curLex.Tag == COMMA {
            if equal {
                comma++
            }else {
                comma--
            }
        }
        if curLex.Tag == VAR {
            if equal{
                depVars = append(depVars, curLex.Image)
            }else{
                vars = append(vars, curLex.Image)
            }
        }
        if (curLex.Tag == ERROR) || 
           (!equal && (curLex.Tag & (VAR | COMMA )) == 0) || 
           ((curLex.Tag == RPAREN) && ((nextLex.Tag & (operations|EQUAL|COMMA|RPAREN|EOF)) == 0)) || 
           ((curLex.Tag == LPAREN) && ((nextLex.Tag & (VAR|NUMBER|MINUS|LPAREN)) == 0)) || 
           ((curLex.Tag == VAR) && ((nextLex.Tag & (operations|EQUAL|COMMA|EOF|RPAREN)) == 0)) ||
           ((curLex.Tag == NUMBER) && ((nextLex.Tag & (operations|COMMA|EOF|RPAREN)) == 0)) ||
           (((curLex.Tag & operations) != 0) && ((nextLex.Tag & (operations|NUMBER|VAR|LPAREN)) == 0)) ||
           ((curLex.Tag == EQUAL) && ((nextLex.Tag & (VAR|NUMBER|LPAREN|MINUS)) == 0)) {
            return nil, nil, false
        }
    }
    if !equal || bracket != 0 || comma != 0 || len(vars) == 0 {
        return nil, nil, false
    }
    return vars, depVars, true
}

func (graph *Graph) dfs(node *Node, seq *[]*Node) bool {
    node.mark = 1
    for _, v := range graph.E[node] {
        if v.mark == 0 {
            ok := graph.dfs(v, seq)
            if !ok {
                return ok
            }
        }else if v.mark == 1 {
            return false 
        }
    }
    node.mark = 2
    *seq = append(*seq, node)
    return true
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    graph := &Graph{make([]*Node, 0), make(map[string]*Node), make(map[*Node][]*Node)}

    for scanner.Scan() {
        formula := scanner.Text()
        if formula == "" {
            break
        }
        vars, depVars, valid := parser(formula)
        if valid {
            node := &Node{formula, vars, depVars, 0} 
            graph.V = append(graph.V, node)
            for _, v := range vars {
                if _, ok := graph.VbyName[v]; ok {
                    fmt.Println("syntax error")
                    return
                }
                graph.VbyName[v] = node
            }
        }else{
            fmt.Println("syntax error")
            return
        }
    }

    for _, v := range graph.V {
        for _, u := range v.depVars {
            alreadyAdd := false
            for _, w := range graph.E[v] {
                if _, ok := graph.VbyName[u]; !ok{
                    fmt.Println("syntax error")
                    return
                }
                if w.formula == graph.VbyName[u].formula {
                    alreadyAdd = true
                    break
                }
            }
            if !alreadyAdd {
                graph.E[v] = append(graph.E[v], graph.VbyName[u])
            }
        }
    }

    seq := make([]*Node, 0)
    for _, v := range graph.V {
        if v.mark != 2 {
            ok := graph.dfs(v, &seq)
            if !ok {
                fmt.Println("cycle")
                return
            }
        }
    }
    for _, node := range seq {
        fmt.Println(node.formula)
    }
}