package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func calc(code []rune) int {
    var operation rune;
    var operands [10]int;
    var count_operands int;
    if len(code) == 2 {
        return int(code[0]) - 48
    }
    for idx := 1; idx < len(code); idx++ {
        if unicode.IsSpace(code[idx]) {
            continue
        }
        if code[idx] == '+' || code[idx] == '-' || code[idx] == '*' {
            operation = code[idx]
        }else if unicode.IsDigit(code[idx]) {
            operands[count_operands] = int(code[idx]) - 48
            count_operands++
        }else if code[idx] == '(' {
            lastIdx := idx + 1
            opened := 0
            for ; lastIdx < len(code); lastIdx++ {
                if code[lastIdx] == '(' {
                    opened++
                }else if opened != 0 && code[lastIdx] == ')'{
                    opened--
                }else if opened == 0 && code[lastIdx] == ')'{
                    break
                } 
            }
            operands[count_operands] = calc(code[idx:lastIdx])
            count_operands++
            idx = lastIdx
        }
    }
    answer := operands[0]
    for i := 1; i < count_operands; i++ {
        if operation == '+' {
            answer += operands[i]
        }else if operation == '-' {
            answer -= operands[i]
        }else {
            answer *= operands[i]
        }
    }
    return answer
}

func main() {
    t, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    text := []rune(t)
    fmt.Println(calc(text))
}
