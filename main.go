package main

import (
	"fmt"

	blockchain "github.com/jonandonigv/blockchain-crypto/block-chain"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("Send 1 BTC to Jon")
	bc.AddBlock("Send 2 more BTC to Jon")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}

}
