package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const SKIP_LIST_LEN = 20

type AssocArray interface {
    Assign(s string, x int)
    Lookup(s string) (x int, exists bool)
}

type SkipListElem struct {
    k string
    v int
    next []*SkipListElem
}

type ElemAVL struct {
    k string
    v, balance int
    parent, left, right *ElemAVL
}

func makeSkipList() AssocArray {
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

func makeAVL() AssocArray {
    return &ElemAVL{
        k: "",
        v: -1,
        parent: nil,
        left: nil,
        right: nil,
    }
}

func Minimum(t *ElemAVL) *ElemAVL{
    x := t
    for x.left != nil {
        x = x.left
    }
    return x
}

func SuccAVL(x *ElemAVL) *ElemAVL {
    if (x.right != nil){
        return Minimum(x.right)
    }
    y := x.parent
    for y != nil && x == y.right {
        x = y
        y = y.parent
    }
    return y
}

func Insert(t *ElemAVL, s string, x int) *ElemAVL {
    y := &ElemAVL{s, x, 0, nil, nil, nil}
    if (t == nil){
        t = y;
    }else{
        r := t;
        for {
			if r.k == s {
				r.v = x;
				return t
			}
            if (s < r.k){
                if (r.left == nil){
                    r.left = y;
                    y.parent = r;
                    break;
                }
                r = r.left;
            }else{
                if (r.right == nil){
                    r.right = y;
                    y.parent = r;
                    break;
                }
                r = r.right;
            }
        }
    }
    return t
}

func ReplaceNode(t, x, y *ElemAVL){
    if (x == t){
        t = y;
        if (y != nil) {
            y.parent = nil
        }
    }else{
        p := x.parent;
        if (y != nil) {
            y.parent = p
        }
        if (p.left == x){
            p.left = y;
        }else{
            p.right = y;
        }
    }
}

func RotateLeft(t, x *ElemAVL) {
    y := x.right
    if y == nil {
        return
    }
    ReplaceNode(t, x, y)
    b := y.left
    if b != nil {
        b.parent = x
    }
    x.right = b
    x.parent = y
    y.left = x
    x.balance--
    if y.balance > 0 {
        x.balance -= y.balance
    }
    y.balance--
    if x.balance < 0 {
        y.balance += x.balance
    }
}

func RotateRight(t *ElemAVL, x *ElemAVL) {
    y := x.left
    if y == nil {
        return
    }
    ReplaceNode(t, x, y)
    b := y.right
    if b != nil {
        b.parent = x
    }
    x.left = b
    x.parent = y
    y.right = x
    x.balance++
    if y.balance < 0 {
        x.balance -= y.balance
    }
    y.balance++
    if x.balance > 0 {
        y.balance += x.balance
    }
}

func (t *ElemAVL) Assign(s string, x int) {
    a := Insert(t, s, x)
    for {
        r := a.parent
        if r == nil {
            break
        }
        if a == r.left {
            r.balance--
            if r.balance == 0 {
                break
            }
            if r.balance == -2 {
                if a.balance == 1 {
                    RotateLeft(t, a)
                }
                RotateRight(t, r)
                break
            }
        }else {
            r.balance++
            if r.balance == 0 {
                break
            }
            if r.balance == 2 {
                if a.balance == -1 {
                    RotateRight(t, a)
                }
                RotateLeft(t, r)
                break
            }
        }
        a = r
    }
}

func Descend(t *ElemAVL, k string) *ElemAVL{
    x := t
    for x != nil && x.k != k {
        if (k < x.k){
            x = x.left
        }else{
            x = x.right
        }
    }
    return x
}

func (t *ElemAVL) Lookup(s string) (x int, exists bool) {
    y := Descend(t, s)
    if (y != nil && y.k == s) {
        return y.v, true
    }
    return 0, false
}

func lex(sentence string, array AssocArray) []int {
    words := strings.Fields(sentence)
    result := make([]int, 0, len(words))
    counter := 1
    for _, word := range words {
        val, exists := array.Lookup(word);
        if exists {
            result = append(result, val)
        } else {
            array.Assign(word, counter)
            result = append(result, counter)
            counter++
        }
    }
    return result
}

func main() {
    sentence := "alpha x1 beta alpha x1 y"
    skipList := makeSkipList()
    resultSkipList := lex(sentence, skipList)
    for _, val := range resultSkipList {
        fmt.Printf("%d ", val)
    }
    fmt.Println()
    
    AVL_Tree := makeAVL()
    resultSkipList = lex(sentence, AVL_Tree)
    for _, val := range resultSkipList {
        fmt.Printf("%d ", val)
    }
}