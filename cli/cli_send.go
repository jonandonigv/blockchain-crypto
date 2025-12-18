package cli

import (
	"fmt"
	"log"

	blockchain "github.com/jonandonigv/blockchain-crypto/block-chain"
)

func (cli *CLI) send(from, to string, amount int) {
	if !blockchain.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}

	if !blockchain.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := blockchain.NewBlockchain(from)
	defer bc.Blocks.Close()

	tx := blockchain.NewUTXOTransaction(from, to, amount, bc)
	bc.AddBlock([]*blockchain.Transaction{tx})
	fmt.Println("Success!")

}
