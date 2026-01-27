package wsserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
	"wallet-app/wallet"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type JSONbalance struct {
	Balance string `json:"balance"`
	Name    string `json:"name"`
}

var myWallet *wallet.Wallet

func StartServer(w *wallet.Wallet) {
	myWallet = w
	http.HandleFunc("/", handleClient)
	go func() {
		addr := myWallet.IPadress + ":" + myWallet.Port
		log.Println("Starting WebSocket server...")
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()
	go ScanNetwork()
}

func ScanNetwork() {
	log.Println("Start scanning network")
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ip := fmt.Sprintf("%s%d", myWallet.IPadressMask, i)
			connection := getConnectionWithWallet(ip)
			if connection == nil {
				return
			}
			defer connection.Close()
			connection.WriteMessage(websocket.TextMessage, []byte("wallet"))
			_, message, err := connection.ReadMessage()
			if err != nil {
				return
			}
			id, err := strconv.Atoi(string(message))
			if err == nil {
				if id == -2 {
					myWallet.ControllerIP = ip
					log.Println("Discovered Controller with id", id)
				} else {
					log.Println("Discovered wallet with id", id)
					myWallet.AddWallet(id, ip)
				}
			}
		}(i)
	}
	wg.Wait()
	log.Println("Finished scanning network")
	updateBalances()
	updateController()
}

func updateBalances() {
	for _, address := range myWallet.GetWallets() {
		connection := getConnectionWithWallet(address)
		defer connection.Close()
		connection.WriteMessage(websocket.TextMessage, []byte("balance"))
		_, message, err := connection.ReadMessage()
		if err != nil {
			fmt.Println(err)
			continue
		}

		var balances map[string]JSONbalance
		if err := json.Unmarshal(message, &balances); err != nil {
			fmt.Println("Error unmarshalling balances:", err)
			continue
		}
		for balanceID, balance := range balances {
			intBalance, _ := strconv.Atoi(balance.Balance)
			intBalanceID, _ := strconv.Atoi(balanceID)
			myWallet.UpdateBalance(intBalanceID, intBalance)
			log.Printf("Updated balance for wallet %d: %d\n", intBalanceID, intBalance)
		}
	}
}

func updateController() {
	if myWallet.ControllerIP != "" {
		conn := getConnectionWithWallet(myWallet.ControllerIP)
		if conn == nil {
			fmt.Println("Failed to connect to", myWallet.ControllerIP)
			return
		}
		defer conn.Close()
		if err := conn.WriteMessage(websocket.TextMessage, []byte("money_balance")); err != nil {
			log.Println("Write error:", err)
			return
		}
		conn.SetReadDeadline(time.Now().Add(3 * time.Second))
		_, reply, _ := conn.ReadMessage()
		if string(reply) == "ok" {
			b := createJSONBalance()
			if err := conn.WriteMessage(websocket.TextMessage, b); err != nil {
				log.Println("Write error:", err)
				return
			}
		}
	}
}

func SendMoney(id int, amount int) {
	_, ok := myWallet.GetWallets()[id]
	if !ok {
		fmt.Println("Wallet not found:", id)
		return
	}
	for walletID, address := range myWallet.GetWallets() {
		conn := getConnectionWithWallet(address)
		if conn == nil {
			fmt.Println("Failed to connect to", address)
			return
		}
		defer conn.Close()
		msg := fmt.Sprintf("send:%d:%d:%d", myWallet.ID, id, amount)
		if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			log.Println("Write error:", err)
			return
		}
		conn.SetReadDeadline(time.Now().Add(3 * time.Second))
		_, reply, _ := conn.ReadMessage()
		if string(reply) == "ok" && walletID == id {
			log.Printf("Sent %d to wallet %d\n", amount, id)
		}
	}
	updateController()
}

func handleClient(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	defer conn.Close()
	_, msg, err := conn.ReadMessage()
	if err != nil {
		return
	}
	message := string(msg)
	ip := strings.Split(r.RemoteAddr, ":")
	log.Printf("Get message from %s: %s\n", ip[0], message)
	switch {
	case message == "wallet":
		if ip[0] == myWallet.ControllerIP {
			updateController()
		} else {
			conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprint(myWallet.ID)))
		}
	case message == "block":
		myWallet.Block = true
	case message == "unblock":
		myWallet.Block = false
	case strings.HasPrefix(message, "send:"):
		parts := strings.Split(message, ":")
		if len(parts) == 4 {
			from, _ := strconv.Atoi(parts[1])
			to, _ := strconv.Atoi(parts[2])
			amount, _ := strconv.Atoi(parts[3])
			myWallet.Balances[from] -= amount
			myWallet.Balances[to] += amount
			conn.WriteMessage(websocket.TextMessage, []byte("ok"))
		}
	case message == "balance":
		b := createJSONBalance()
		conn.WriteMessage(websocket.TextMessage, b)
	default:
		conn.WriteMessage(websocket.TextMessage, []byte("Unknown command"))
	}
}

func getConnectionWithWallet(address string) *websocket.Conn {
	ip := fmt.Sprintf("ws://%s:3000/ws", address)
	connection, _, err := websocket.DefaultDialer.Dial(ip, nil)
	if err != nil {
		return nil
	}
	return connection
}

func createJSONBalance() []byte {
	var balances = make(map[int]JSONbalance)
	for id, balance := range myWallet.Balances {
		strBalance := fmt.Sprintf("%d", balance)
		name := ""
		if id == myWallet.ID {
			name = myWallet.UserName
		}
		balances[id] = JSONbalance{Balance: strBalance, Name: name}
	}
	b, _ := json.Marshal(balances)
	return b
}
