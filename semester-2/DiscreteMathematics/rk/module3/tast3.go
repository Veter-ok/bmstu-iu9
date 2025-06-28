package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PairState struct {
    q0, q1 int
}

type DoubleMealy struct {
    n, m, q0 int
    states []*PairState
    delta map[*PairState]map[string]*PairState
    phi map[*PairState]map[string]string
}

type Mealy struct {
    n, m, q0 int
    delta map[int]map[string]int
    phi map[int]map[string]string
}

func CombainMealy(mealy1, mealy2 *Mealy) *DoubleMealy {
    n := mealy1.n * mealy2.n
    m := mealy1.m
    delta := make(map[*PairState]map[string]*PairState)
    phi := make(map[*PairState]map[string]string)
    states := make([]*PairState, 0)
    checkPair := make(map[int]map[int]*PairState)
    for q1 := 0; q1 < mealy1.n; q1++ {
        for q2 := 0; q2 < mealy2.n; q2++ {
            for s := 0; s < m; s++ {
                pairState1 := &PairState{q1, q2}
                states = append(states, pairState1)
                checkPair[q1] = make(map[int]*PairState)
                checkPair[q1][q2] = pairState1
                
                delta[pairState1] = make(map[string]*PairState)
                phi[pairState1] = make(map[string]string)

                symbolAfterM1 := mealy1.phi[q1][string('a' + s)]
                stateAfterM1 := mealy1.delta[q2][string('a' + s)]

                phi[pairState1][string('a' + s)] = mealy2.phi[stateAfterM1][symbolAfterM1]
                pairState2 := &PairState{stateAfterM1, mealy2.delta[stateAfterM1][symbolAfterM1]}
                if _, ok := checkPair[pairState1.q0][pairState1.q1]; ok {
                    pairState2 = checkPair[pairState1.q0][pairState1.q1]
                }else{
                    states = append(states, pairState2)
                    checkPair[pairState2.q0] = make(map[int]*PairState)
                    checkPair[pairState2.q0][pairState2.q1] = pairState2
                }
                
                delta[pairState1][string('a' + s)] = pairState2
            }
        }
    }
    return &DoubleMealy{n, m, 0, states, delta, phi}
} 

func readMealy(reader *bufio.Reader) *Mealy {
    var n, m, q0 int
    fmt.Fscanf(reader, "%d %d %d\n", &n, &m, &q0)

    delta := make(map[int]map[string]int, n)
    for i := 0; i < n; i++ {
        line, _ := reader.ReadString('\n')
        parts := strings.Fields(line)
        delta[i] = make(map[string]int, m)
        for j := 0; j < m; j++ {
            delta[i][string('a' + j)], _ = strconv.Atoi(parts[j])
        }
    }

    phi := make(map[int]map[string]string, n)
    for i := 0; i < n; i++ {
        line, _ := reader.ReadString('\n')
        parts := strings.Fields(line)
        phi[i] = make(map[string]string, m)
        for j := 0; j < m; j++ {
            phi[i][string('a' + j)] = parts[j]
        }
    }

    return &Mealy{n, m, q0, delta, phi}
}

func (mealy DoubleMealy) dfs(q0 int) {

}

func main(){
    reader := bufio.NewReader(os.Stdin)

    mealy1 := readMealy(reader)
    mealy2 := readMealy(reader)

    doubleMealy := CombainMealy(mealy1, mealy2)

    fmt.Println(doubleMealy.n, doubleMealy.m, doubleMealy.q0)
    for _, state := range doubleMealy.states {
        for j := 0; j < doubleMealy.m; j++ {
            symbol := string('a' + j)
            fmt.Printf("%d ", doubleMealy.delta[state][symbol])
        }
        fmt.Println()
    }

    for _, state := range doubleMealy.states {
        for j := 0; j < doubleMealy.m; j++ {
            symbol := string('a' + j)
            fmt.Printf("%s ", doubleMealy.phi[state][symbol])
        }
        fmt.Println()
    }
}