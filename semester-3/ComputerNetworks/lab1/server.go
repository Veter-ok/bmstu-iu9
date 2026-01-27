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
	l, err := net.Listen("tcp4", ":3000")
	if err != nil {
		fmt.Println(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleClient(c)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	msg := Message{}

	for {
		n, _ := conn.Read(buffer)
		json.Unmarshal(buffer[:n], &msg)
		fmt.Println("Get data")
		fmt.Println(msg)
		intersect, x, y := calculateIntersection(msg.X1, msg.Y1, msg.X2, msg.Y2, msg.X3, msg.Y3, msg.X4, msg.Y4)
		response := Response{intersect, x, y}
		// responseData, _ := json.Marshal(response)
		if response.Intersect {
			conn.Write([]byte("Intersect"))
		} else {
			conn.Write([]byte("Not Intersect"))
		}
		fmt.Println("Sent response")
		break
	}
}

func calculateIntersection(x1, y1, x2, y2, x3, y3, x4, y4 float32) (bool, float32, float32) {
	denom := (y4-y3)*(x2-x1) - (x4-x3)*(y2-y1)
	if denom == 0 {
		return false, 0, 0
	}
	ua := ((x4-x3)*(y1-y3) - (y4-y3)*(x1-x3)) / denom
	return true, x1 + ua*(x2-x1), y1 + ua*(y2-y1)
}
