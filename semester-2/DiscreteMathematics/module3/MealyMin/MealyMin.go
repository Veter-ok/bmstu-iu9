package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type State struct {
    parent *State
    depth, i  int
}

type MealyMachine struct {
    Q []*State
    X []string
    Delta map[*State]map[string]*State
    Phi map[*State]map[string]string
}

func Find(q *State) *State {
    if q.parent != q {
        q.parent = Find(q.parent)
    }
    return q.parent
}

func Union(q1, q2 *State) {
    root1 := Find(q1)
    root2 := Find(q2)
    if root1 == root2 {
        return
    }
    if root1.depth < root2.depth {
        root1.parent = root2
    } else {
        root2.parent = root1
        if root1.depth == root2.depth {
            root1.depth++
        }
    }
}

func Contain(Q []*State, q *State) bool {
    for _, q2 := range Q {
        if q2.i == q.i {
            return true
        }
    }
    return false
}

func split1(machine MealyMachine) (int, []*State) {
    m := len(machine.Q)
    for _, q := range machine.Q {
        q.parent = q
        q.depth = 0
    }
    for i := 0; i < len(machine.Q); i++ {
        for j := i + 1; j < len(machine.Q); j++ {
            q1 := machine.Q[i]
            q2 := machine.Q[j]
            if Find(q1) != Find(q2) {
                eq := true
                for _, x := range machine.X {
                    if machine.Phi[q1][x] != machine.Phi[q2][x] {
                        eq = false
                        break
                    }
                }
                if eq {
                    Union(q1, q2)
                    m = m - 1
                }
            }
        }
    }
    pi := make([]*State, len(machine.Q))
    for _, q := range machine.Q {
        pi[q.i] = Find(q)
    }
    return m, pi
}

func split(machine MealyMachine, pi []*State) (int, []*State) {
    m := len(machine.Q)
    for _, q := range machine.Q {
        q.parent = q
        q.depth = 0
    }

    for i := 0; i < len(machine.Q); i++ {
        for j := i + 1; j < len(machine.Q); j++ {
            q1 := machine.Q[i]
            q2 := machine.Q[j]
            if pi[q1.i] == pi[q2.i] && Find(q1) != Find(q2) {
                eq := true
                for _, x := range machine.X {
                    w1, w2 := machine.Delta[q1][x], machine.Delta[q2][x]
                    if pi[w1.i] != pi[w2.i] {
                        eq = false
                        break
                    }
                }
                if eq {
                    Union(q1, q2)
                    m--
                }
            }
        }
    }

    newPi := make([]*State, len(machine.Q))
    for _, q := range machine.Q {
        newPi[q.i] = Find(q)
    }
    return m, newPi
}

func AufenkampHohn(A MealyMachine) MealyMachine {
    m, pi := split1(A)
    for {
        m2, pi2 := split(A, pi)
        pi = pi2
        if m == m2 {
            break
        }
        m = m2
    }
    new_Q := make([]*State, 0)
    new_delta := make(map[*State]map[string]*State)
    new_phi := make(map[*State]map[string]string)

    for _, q := range A.Q {
        qq := pi[q.i]
        if !Contain(new_Q, qq) {
            new_Q = append(new_Q, qq)
            new_delta[qq] = make(map[string]*State)
            new_phi[qq] = make(map[string]string)
            for _, x := range A.X {
                new_delta[qq][x] = pi[A.Delta[q][x].i]
                new_phi[qq][x] = A.Phi[q][x]
            }
        }
    }

    return MealyMachine{new_Q, A.X, new_delta, new_phi}
}

func dfs(machine *MealyMachine, q0 *State, visited map[*State]bool, order *[]*State) {
    if visited[q0] {
        return
    }
    visited[q0] = true
    *order = append(*order, q0)
    for _, x := range machine.X {
        q := machine.Delta[q0][x]
        if !visited[q] {
            dfs(machine, q, visited, order)
        }
    }
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    n, _ := strconv.Atoi(scanner.Text())
    scanner.Scan()
    m, _ := strconv.Atoi(scanner.Text())
    scanner.Scan()
    q0, _ := strconv.Atoi(scanner.Text())

    Q := make([]*State, n)
    for i := 0; i < n; i++ {
        Q[i] = &State{nil, 0, i}
    }

    X := make([]string, m)
    for i := 0; i < m; i++ {
        X[i] = string('a' + i)
    }

    delta := make(map[*State]map[string]*State)
    for i := 0; i < n; i++ {
        scanner.Scan()
        parts := strings.Fields(scanner.Text())
        delta[Q[i]] = make(map[string]*State)
        for j := 0; j < m; j++ {
            idx, _ := strconv.Atoi(parts[j])
            delta[Q[i]][X[j]] = Q[idx]
        }
    }

    phi := make(map[*State]map[string]string)
    for i := 0; i < n; i++ {
        scanner.Scan()
        parts := strings.Fields(scanner.Text())
        phi[Q[i]] = make(map[string]string)
        for j := 0; j < m; j++ {
            phi[Q[i]][X[j]] = parts[j]
        }
    }

    machine := MealyMachine{Q, X, delta, phi}
    new_A := AufenkampHohn(machine)

    q0State := new_A.Q[0]
    for _, q := range Q {
        if q.i == q0 {
            q0State = q.parent
            break
        }
    }

    visited := make(map[*State]bool, len(new_A.Q))
    order := make([]*State, 0)
    dfs(&new_A, q0State, visited, &order)

    for i, q1 := range order {
        q1.i = i
    }

    fmt.Println("digraph {\n    rankdir = LR")
    for _, q1 := range order {
        for _, x := range new_A.X {
            q2 := new_A.Delta[q1][x]
            fmt.Printf("    %d -> %d [label = \"%s(%s)\"]\n", q1.i, q2.i, x, new_A.Phi[q1][x])
        }
    }
    fmt.Println("}")
}