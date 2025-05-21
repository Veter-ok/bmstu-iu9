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

func dfs(machine *MealyMachine, q0 *State, visited map[*State]bool, order *([]*State)) {
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

func new_numeration(machine *MealyMachine, Q []*State, q0 int) []*State {
    q0State := machine.Q[0]
    for _, q := range Q {
        if q.i == q0 {
            q0State = q.parent
            break
        }
    }
    visited := make(map[*State]bool, len(machine.Q))
    order := make([]*State, 0)
    dfs(machine, q0State, visited, &order)
    for i, q := range order {
        q.i = i
    }
    return order
}

func isEqual(machine1, machine2 MealyMachine) bool {
    if len(machine1.Q) != len(machine2.Q) {
        return false
    }
    for i := 0; i < len(machine1.Q); i++ {
        q1, q2 := machine1.Q[i], machine2.Q[i]
        if q1.i != q2.i {
            return false
        }
        for _, x := range machine1.X {
            if (machine1.Phi[q1][x] != machine2.Phi[q2][x]) || 
               (machine1.Delta[q1][x].i != machine2.Delta[q2][x].i) {
                return false
            }
        }
    }
    return true
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    n1, _ := strconv.Atoi(scanner.Text())
    scanner.Scan()
    m1, _ := strconv.Atoi(scanner.Text())
    scanner.Scan()
    q0_1, _ := strconv.Atoi(scanner.Text())
    Q1 := make([]*State, n1)
    for i := 0; i < n1; i++ {
        Q1[i] = &State{nil, 0, i}
    }
    X1 := make([]string, m1)
    for i := 0; i < m1; i++ {
        X1[i] = string('a' + i)
    }
    delta1 := make(map[*State]map[string]*State)
    for i := 0; i < n1; i++ {
        scanner.Scan()
        parts := strings.Fields(scanner.Text())
        delta1[Q1[i]] = make(map[string]*State)
        for j := 0; j < m1; j++ {
            idx, _ := strconv.Atoi(parts[j])
            delta1[Q1[i]][X1[j]] = Q1[idx]
        }
    }
    phi1 := make(map[*State]map[string]string)
    for i := 0; i < n1; i++ {
        scanner.Scan()
        parts := strings.Fields(scanner.Text())
        phi1[Q1[i]] = make(map[string]string)
        for j := 0; j < m1; j++ {
            phi1[Q1[i]][X1[j]] = parts[j]
        }
    }

    scanner.Scan()
    n2, _ := strconv.Atoi(scanner.Text())
    scanner.Scan()
    m2, _ := strconv.Atoi(scanner.Text())
    scanner.Scan()
    q0_2, _ := strconv.Atoi(scanner.Text())
    Q2 := make([]*State, n2)
    for i := 0; i < n2; i++ {
        Q2[i] = &State{nil, 0, i}
    }
    X2 := make([]string, m2)
    for i := 0; i < m2; i++ {
        X2[i] = string('a' + i)
    }
    delta2 := make(map[*State]map[string]*State)
    for i := 0; i < n2; i++ {
        scanner.Scan()
        parts := strings.Fields(scanner.Text())
        delta2[Q2[i]] = make(map[string]*State)
        for j := 0; j < m2; j++ {
            idx, _ := strconv.Atoi(parts[j])
            delta2[Q2[i]][X2[j]] = Q2[idx]
        }
    }
    phi2 := make(map[*State]map[string]string)
    for i := 0; i < n2; i++ {
        scanner.Scan()
        parts := strings.Fields(scanner.Text())
        phi2[Q2[i]] = make(map[string]string)
        for j := 0; j < m2; j++ {
            phi2[Q2[i]][X2[j]] = parts[j]
        }
    }

    machine1 := MealyMachine{Q1, X1, delta1, phi1}
    newA1 := AufenkampHohn(machine1)
    newA1.Q = new_numeration(&newA1, Q1, q0_1)

    machine2 := MealyMachine{Q2, X2, delta2, phi2}
    newA2 := AufenkampHohn(machine2)
    newA2.Q = new_numeration(&newA2, Q2, q0_2)

    if isEqual(newA1, newA2) {
        fmt.Println("EQUAL")
    }else{
        fmt.Println("NOT EQUAL")
    }
}