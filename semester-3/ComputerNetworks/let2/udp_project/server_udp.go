package main

import (
	"fmt"
	"math/rand/v2"
	"net"
	"strings"
)

var NUM string
var collect map[int]bool

func main() {
	addr, err := net.ResolveUDPAddr("udp", "10.37.196.182:4000")
	if err != nil {
		fmt.Println("Ошибка разрешения адреса:", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
		return
	}
	defer conn.Close()

	fmt.Println("UDP сервер запущен на порту 4000")
	fmt.Println("Ожидание сообщений...")

	collect = make(map[int]bool)
	NUM = guessNumber()

	buffer := make([]byte, 1024)

	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Ошибка чтения:", err)
			continue
		}

		message := strings.TrimSpace(string(buffer[:n]))
		fmt.Printf("Получено сообщение от %s: %s\n", clientAddr, message)

		go handleClient(conn, clientAddr, message)
	}
}

func handleClient(conn *net.UDPConn, clientAddr *net.UDPAddr, message string) {
	cow, bull := analyzeMessage(message, NUM)

	var response string
	if bull == 4 {
		response = "Поздравляем! Вы угадали число!\n"
		NUM = guessNumber()
	} else {
		response = fmt.Sprintf("Коровы: %d, Быки: %d\n", cow, bull)
	}

	_, err := conn.WriteToUDP([]byte(response), clientAddr)
	if err != nil {
		fmt.Println("Ошибка отправки:", err)
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
	ans := ""
	collect = make(map[int]bool)
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
	fmt.Println("Новое число:", ans)
	return ans
}
