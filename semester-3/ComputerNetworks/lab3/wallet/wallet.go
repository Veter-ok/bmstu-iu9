package wallet

import "log"

type Wallet struct {
	ID           int
	UserName     string
	IPadressMask string
	ControllerIP string
	IPadress     string
	Port         string
	Balance      int
	Block        bool
	wallets      map[int]string
	Balances     map[int]int
}

func NewWallet(id int, userName, IPadressMask, IPadress, port string, balance int) *Wallet {
	log.Println("Initializing wallet")
	return &Wallet{
		ID:           id,
		UserName:     userName,
		IPadressMask: IPadressMask,
		IPadress:     IPadress,
		ControllerIP: "",
		Port:         port,
		Balance:      balance,
		Block:        false,
		wallets:      make(map[int]string),
		Balances:     make(map[int]int),
	}
}

func (w *Wallet) GetWallets() map[int]string {
	return w.wallets
}

func (w *Wallet) AddWallet(id int, address string) {
	w.wallets[id] = address
	w.Balances[id] = 10
}

func (w *Wallet) UpdateBalance(id int, amount int) {
	w.Balances[id] = amount
}
