package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"net"
	"strings"
)

var NUM string
var collect map[int]bool

func main() {
	listener, err := net.Listen("tcp", "10.37.196.182:4000")
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Сервер запущен на порту 4000")
	fmt.Println("Ожидание подключений...")

	collect = make(map[int]bool)
	NUM = guessNumber()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка подключения:", err)
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Новое подключение")
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Клиент отключился")
			return
		}
		fmt.Printf("Клиент ввёл: %s", message)
		cow, bull := analyzeMessage(strings.TrimSpace(message), NUM)
		if bull == 4 {
			conn.Write([]byte("Поздравляем! Вы угадали число!\n"))
			NUM = guessNumber()
		} else {
			response := fmt.Sprintf("Коровы: %d, Быки: %d\n", cow, bull)
			conn.Write([]byte(response))
		}
	}
}

func analyzeMessage(message string, number string) (int, int) {
	cow, bull := 0, 0
	for i := 0; i < 4; i++ {
		if strings.Contains(message, string(number[i])) {
			if string(number[i]) == string(message[i]) {
				bull++
			} else {
				cow++
			}
		}
	}
	return cow, bull
}

func guessNumber() string {
	collect = make(map[int]bool)
	ans := ""
	for i := 0; i < 4; i++ {
		d := rand.IntN(10)
		for {
			if _, ok := collect[d]; !ok {
				break
			}
			d = rand.IntN(10)
		}
		ans += string('0' + d)
		collect[d] = true
	}
	fmt.Println(ans)
	return ans
}
