package cli

import (
	"fmt"
	"log"

	blockchain "github.com/jonandonigv/blockchain-crypto/block-chain"
)

func (cli *CLI) listAddresses() {
	wallets, err := blockchain.NewWallets()
	if err != nil {
		log.Panic(err)
	}

	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}
