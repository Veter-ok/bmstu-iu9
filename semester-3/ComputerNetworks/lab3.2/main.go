package main

import (
	"bufio"
	"fmt"
	"os"
	"peer-lab/peer"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	fmt.Println("Введите имя узла: ")
	name, _ := scanner.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("Введите свой порт: ")
	ipAdress, _ := scanner.ReadString('\n')
	ipAdress = strings.TrimSpace(ipAdress)

	fmt.Println("Введите порт родителя: ")
	ipParent, _ := scanner.ReadString('\n')
	ipParent = strings.TrimSpace(ipParent)

	MyPeer := peer.CreatePeer(name, ipAdress, ipParent)
	go peer.StartPeer(MyPeer)

	fmt.Println("1. Имя родителя\n2. Все потомки\n3. Завершение")
	for {
		text, err := scanner.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		commandId, _ := strconv.Atoi(strings.TrimSpace(text))
		switch commandId {
		case 1:
			if MyPeer.IPParent == "" {
				fmt.Println("Главный peer")
			} else {
				name := MyPeer.GetNameFromParent()
				fmt.Printf("Имя родителя: %s\n", name)
			}
		case 2:
			if MyPeer.IPParent == "" {
				fmt.Println("Главный peer")
			} else {
				names := MyPeer.GetTotalName()
				names = append([]string{MyPeer.PeerName}, names...)
				for i, name := range names {
					fmt.Println(name)
					if i != len(names)-1 {
						fmt.Println(" ↑")

					}
				}
			}
		case 3:
			if MyPeer.IPParent == "" {
				fmt.Println("Главный peer")
			} else {
				MyPeer.Exit()
			}
			return
		}
	}
}
