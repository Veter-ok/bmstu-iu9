package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Подключаемся к серверу
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("Не удалось подключиться к серверу:", err)
		return
	}
	defer conn.Close() // Закрываем соединение при выходе

	fmt.Println("Подключено к серверу")
	fmt.Println("Вводите сообщения (Ctrl+C для выхода):")

	// Го-рутина для чтения ответов от сервера
	go readFromServer(conn)

	// Создаем читатель для ввода с клавиатуры
	scanner := bufio.NewScanner(os.Stdin)

	// Бесконечный цикл чтения ввода пользователя
	for scanner.Scan() {
		text := scanner.Text()

		// Отправляем сообщение серверу
		conn.Write([]byte(text + "\n"))
	}
}

func readFromServer(conn net.Conn) {

	// Создаем читатель для входящих данных от сервера
	reader := bufio.NewReader(conn)

	// Бесконечный цикл чтения сообщений от сервера
	for {
		// Читаем строку от сервера
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Соединение с сервером разорвано")
			os.Exit(0) // Выходим из программы
		}

		// Выводим сообщение от сервера
		fmt.Print("Сервер: ", message)
	}
}
