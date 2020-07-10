package main

import (
	"fmt"
	"go-blockchain/pkg/blockchain"
)

//TODO: Make this an http api to allow users to interact with the blockchain.
func main() {
	bc := blockchain.New()
	bc.AddBlock("testing")

	for _, block := range bc.Blocks {
		fmt.Println(block)
	}
}
