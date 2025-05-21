package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"unicode"
)

type Tag int

const SKIP_LIST_LEN = 20

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
)

type Lexem struct {
    Tag
    Image string
}

type Parser struct {
    lexems chan Lexem
    current Lexem
}

type SkipListElem struct {
    k string
    v int
    next []*SkipListElem
}

func makeSkipList() *SkipListElem {
    return &SkipListElem{
        k: "",
        v: -1,
        next: make([]*SkipListElem, SKIP_LIST_LEN),
    }
}

func Skip(l *SkipListElem, k string, p []*SkipListElem) {
    for i := SKIP_LIST_LEN - 1; i >= 0; i-- {
        for l.next[i] != nil && l.next[i].k < k {
            l = l.next[i]
        }
        p[i] = l
    }
}

func SuccSkipList(l *SkipListElem) *SkipListElem {
    return l.next[0];
}

func (l *SkipListElem) Lookup(s string) (x int, exists bool) {
    p := make([]*SkipListElem, SKIP_LIST_LEN)
    Skip(l, s, p)
    t := SuccSkipList(p[0])
    if t != nil && t.k == s {
        return t.v, true
    }
    return 0, false
}

func (l *SkipListElem) Assign(s string, x int) {
    p := make([]*SkipListElem, SKIP_LIST_LEN)
    Skip(l, s, p)
    if p[0].next[0] != nil && p[0].next[0].k == s {
        p[0].next[0].v = x
        return
    }
    t := &SkipListElem{s, x, make([]*SkipListElem, SKIP_LIST_LEN)}
    r := rand.Int31n(1000) * 2
    for i := 0; i < SKIP_LIST_LEN && r%2 == 0; i++ {
        t.next[i] = p[i].next[i]
        p[i].next[i] = t
        r /= 2
    }
}

func lexer(expr string, lexems chan Lexem) {
    for i := 0; i < len(expr); i++ {
        switch {
        case unicode.IsSpace(rune(expr[i])):
            continue;
        case expr[i] == '+':
            lexems <- Lexem{PLUS, string(expr[i])}
        case expr[i] == '-':
            lexems <- Lexem{MINUS, string(expr[i])}
        case expr[i] == '*':
            lexems <- Lexem{MUL, string(expr[i])}
        case expr[i] == '/':
            lexems <- Lexem{DIV, string(expr[i])}
        case expr[i] == '(':
            lexems <- Lexem{LPAREN, string(expr[i])}
        case expr[i] == ')':
            lexems <- Lexem{RPAREN, string(expr[i])}
        case unicode.IsDigit(rune(expr[i])):
            j := i 
            for j < len(expr) && unicode.IsDigit(rune(expr[j])) {
                j++
            }
            lexems <- Lexem{NUMBER, expr[i:j]}
            i = j - 1
        case unicode.IsLetter(rune(expr[i])):
            j := i 
            for j < len(expr) && (unicode.IsDigit(rune(expr[j])) || unicode.IsLetter(rune(expr[j]))) {
                j++
            }
            lexems <- Lexem{VAR, expr[i:j]}
            i = j - 1;
        default:
            lexems <- Lexem{ERROR, string(expr[i])}
        }
    }
    close(lexems)
}

func makeParser(lexems chan Lexem) *Parser {
    return &Parser{lexems: lexems}
}

func (parser *Parser) next() {
    parser.current = <-parser.lexems
}

// <E>  ::= <T> <E'>.
func (parser *Parser) parseE(skipList *SkipListElem) (int, error) {
    val, err := parser.parseT(skipList)
    if err != nil {
        return 0, err
    }
    if parser.current.Tag == ERROR {
        return 0, fmt.Errorf("error")
    }
    return parser.parseEe(val, skipList)
}

// <E'> ::= + <T> <E'> | - <T> <E'> | .
func (parser *Parser) parseEe(val int, skipList *SkipListElem) (int, error) {
    if parser.current.Tag & (PLUS | MINUS) != 0 {
        op := parser.current
        parser.next()
        val2, err := parser.parseT(skipList)
        if err != nil {
            return 0, err
        }
        if op.Tag == PLUS {
            val += val2
        }else {
            val -= val2
        }
        return parser.parseEe(val, skipList)
    }
    return val, nil
}

// <T>  ::= <F> <T'>.
func (parser *Parser) parseT(skipList *SkipListElem) (int, error) {
    val, err := parser.parseF(skipList)
    if err != nil {
        return 0, err
    }
    return parser.parseTt(val, skipList)
}

// <T'> ::= * <F> <T'> | / <F> <T'> | .
func (parser *Parser) parseTt(val int, skipList *SkipListElem) (int, error) {
    if parser.current.Tag == ERROR {
        return 0, fmt.Errorf("error")
    }
    if parser.current.Tag & (MUL | DIV) != 0 {
        op := parser.current
        parser.next()
        val2, err := parser.parseF(skipList)
        if err != nil {
            return 0, err
        }
        if op.Tag == MUL {
            val *= val2
        }else {
            val /= val2
        }
        return parser.parseTt(val, skipList)
    }
    return val, nil
}

// <F>  ::= <number> | <var> | ( <E> ) | - <F>.
func (parser *Parser) parseF(skipList *SkipListElem) (int, error) {
    switch parser.current.Tag {
    case NUMBER:
        val, _ := strconv.Atoi(parser.current.Image)
        parser.next()
        return val, nil
    case VAR:
        val, exists := skipList.Lookup(parser.current.Image)
        if exists {
            parser.next()
            return val, nil
        }
        fmt.Scan(&val)
        skipList.Assign(parser.current.Image, val)
        parser.next()
        return val, nil
    case LPAREN:
        parser.next()
        val, err := parser.parseE(skipList)
        if err != nil {
            return 0, err
        }
        parser.next()
        return val, nil
    case MINUS:
        parser.next()
        val, err := parser.parseF(skipList)
        if err != nil {
            return 0, err
        }
        return -val, nil
    default:
        return 0, fmt.Errorf("error")
    }
}

func main(){
    sentence := os.Args[1]
    skipList := makeSkipList()
    result := make(chan Lexem)
    go lexer(sentence, result)
    parser := makeParser(result)
    parser.next()
    val, err := parser.parseE(skipList)
    if err != nil {
        fmt.Println(err)
    }else if parser.current.Tag != 0 {
        fmt.Println("error")
    }else {
        fmt.Println(val)
    }
}