package cli

import (
	"fmt"
	"log"

	blockchain "github.com/jonandonigv/blockchain-crypto/block-chain"
)

func (cli *CLI) createBlockchain(address string) {
	if !blockchain.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}

	bc := blockchain.CreateBlockchain(address)
	bc.Blocks.Close()
	fmt.Println("Done!")
}
