package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

type FireDB struct {
	*db.Client
}

var fireDB FireDB

func (db *FireDB) Connect() error {
	ctx := context.Background()
	opt := option.WithCredentialsFile("bmstu-sem2-firebase-adminsdk-fbsvc-159e4e53ed.json")
	config := &firebase.Config{DatabaseURL: "https://bmstu-sem2-default-rtdb.firebaseio.com/"}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		return fmt.Errorf("error initializing app: %v", err)
	}
	client, err := app.Database(ctx)
	if err != nil {
		return fmt.Errorf("error initializing database: %v", err)
	}
	db.Client = client
	return nil
}

func FirebaseDB() *FireDB {
	return &fireDB
}

func main() {
	ctx := context.Background()

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/6f1c6a8fd16b4d9ba6d2565dc3c47920")
	if err != nil {
		log.Fatalln(err)
	}

	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(header.Number.Int64())
	block, err := client.BlockByNumber(ctx, blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	err = FirebaseDB().Connect()
	if err != nil {
		log.Println(err)
		return
	}

	var txData map[string]interface{}

	for _, tx := range block.Transactions() {
		txData = map[string]interface{}{
			"chainId":  tx.ChainId().String(),
			"hash":     tx.Hash().Hex(),
			"value":    tx.Value().String(),
			"cost":     tx.Cost().String(),
			"to":       tx.To().String(),
			"gas":      tx.Gas(),
			"gasPrice": tx.GasPrice().String(),
			"block":    block.Number().Uint64(),
		}
		break
	}

	ref := fireDB.NewRef("lastTransaction")

	if err := ref.Set(ctx, txData); err != nil {
		log.Fatalf("Firebase write error: %v", err)
	}
	fmt.Println(txData)
}
