package main

import (
	"fmt"
	"wallet-app/wallet"
	wsserver "wallet-app/wsServer"
)

func main() {
	wallet := wallet.NewWallet(3, "Rodion", "172.20.10.", "172.20.10.2", "3000", 10)

	go wsserver.StartServer(wallet)
	for {
		fmt.Println("1. Show wallets")
		fmt.Println("2. Show balances")
		fmt.Println("3. Send money")
		fmt.Println("4. Exit")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Wallets:")
			for id, address := range wallet.GetWallets() {
				fmt.Printf("ID: %d, Address: %s\n", id, address)
			}
		case 2:
			fmt.Println("Balances:")
			for id, balance := range wallet.Balances {
				fmt.Printf("ID: %d, Balance: %d\n", id, balance)
			}
		case 3:
			var id, amount int
			fmt.Print("Enter wallet ID to send money to: ")
			fmt.Scan(&id)
			fmt.Print("Enter amount to send: ")
			fmt.Scan(&amount)
			wsserver.SendMoney(id, amount)
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
