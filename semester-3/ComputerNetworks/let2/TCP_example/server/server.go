package main

import (
    "bufio"
    "fmt"
    "net"
)

func main() {
    // Слушаем порт 9999 для входящих подключений
    listener, err := net.Listen("tcp", ":9999")
    if err != nil {
        fmt.Println("Ошибка запуска сервера:", err)
        return
    }
    defer listener.Close() // Закрываем слушатель при выходе
    
    fmt.Println("Сервер запущен на порту 9999")
    fmt.Println("Ожидание подключений...")
    
    // Бесконечный цикл принятия подключений
    for {
        // Принимаем новое подключение
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Ошибка подключения:", err)
            continue
        }
        
        // Обрабатываем каждое подключение в отдельной горутине
        go handleClient(conn)
    }
}

func handleClient(conn net.Conn) {
    defer conn.Close() // Закрываем соединение при выходе из функции
    
    fmt.Println("Новое подключение")
    
    // Создаем читатель для входящих данных
    reader := bufio.NewReader(conn)
    
    // Бесконечный цикл чтения сообщений от клиента
    for {
        // Читаем строку до символа новой строки
        message, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Клиент отключился")
            return
        }
        
        // Выводим полученное сообщение
        fmt.Print("Получено: ", message)
        
        // Отправляем ответ клиенту
        response :=  message
        conn.Write([]byte(response))
    }
}