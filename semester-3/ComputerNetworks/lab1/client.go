package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type Message struct {
	X1, Y1, X2, Y2 float32
	X3, Y3, X4, Y4 float32
}

type Response struct {
	Intersect bool
	X, Y      float32
}

func main() {
	var x1, y1, x2, y2, x3, y3, x4, y4 float32
	fmt.Print("Line 1: ")
	fmt.Scanf("%f %f %f %f", &x1, &y1, &x2, &y2)
	fmt.Print("Line 2: ")
	fmt.Scanf("%f %f %f %f", &x3, &y3, &x4, &y4)

	msg := Message{x1, y1, x2, y2, x3, y3, x4, y4}
	msgEncode, err := json.Marshal(msg)

	conn, err := net.Dial("tcp", "185.102.139.161:5050")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer conn.Close()
	_, err = conn.Write(msgEncode)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)

	var response Response
	json.Unmarshal(buffer[:n], &response)
	if response.Intersect {
		fmt.Printf("(%.2f, %.2f)\n", response.X, response.Y)
	} else {
		fmt.Println("Lines do not intersect")
	}
}
