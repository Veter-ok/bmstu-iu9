package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func econom_calc(code string) int {
    var answer int
    for idx := strings.Index(code, ")"); idx > 0; idx, answer = strings.Index(code, ")"), answer + 1 {
        operand := code[idx - 4:idx+1]
        mark := string(rune('?' + idx))
        code = strings.Replace(code, operand, mark, -1)
    }
    return answer
}

func main(){
    text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    fmt.Println(econom_calc(text))
}