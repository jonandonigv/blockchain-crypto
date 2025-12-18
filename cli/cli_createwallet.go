package cli

import (
	"fmt"

	blockchain "github.com/jonandonigv/blockchain-crypto/block-chain"
)

func (cli *CLI) createWallet() {
	wallets, _ := blockchain.NewWallets()
	address := wallets.CreateWallet()
	wallets.SaveToFile()

	fmt.Println("Your new address &s\n", address)
}
