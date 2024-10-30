package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://localhost:8545")
	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client:", err)
		return
	}
	fmt.Println("We have a connection to the Ethereum client:", client)
}