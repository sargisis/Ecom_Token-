package main

import (
	"context"
	"fmt"
	"log"
	
	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/types"
	"github.com/mr-tron/base58"
)

func main() {
	
	c := client.NewClient("https://api.devnet.solana.com")

	
	payer := types.NewAccount()
	fmt.Println("Testing private key:", base58.Encode(payer.PrivateKey))
	fmt.Println("Tesing public key:", payer.PublicKey.ToBase58())

	
	fmt.Println("1. go to  https://faucet.solana.com")
	fmt.Println("2. Your public key:", payer.PublicKey.ToBase58())
	fmt.Println("3. Запросите тестовые SOL")
	fmt.Println("4. Нажмите Enter, чтобы продолжить...")
	fmt.Scanln() 

	
	balance, err := c.GetBalance(context.Background(), payer.PublicKey.ToBase58())
	if err != nil {
		log.Fatal("Error:", err)
	}
	if balance == 0 {
		log.Fatal("Balance is 0.")
	}
	fmt.Printf("Balence: %.4f SOL\n", float64(balance)/1e9)

	
	fmt.Println("Balance is getting")
}