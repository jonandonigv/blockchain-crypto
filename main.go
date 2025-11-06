package main

import (
	blockchain "github.com/jonandonigv/blockchain-crypto/block-chain"
	"github.com/jonandonigv/blockchain-crypto/cli"
)

func main() {
	bc := blockchain.NewBlockchain()
	defer bc.Blocks.Close()

	cli := cli.CLI{Bc: bc}
	cli.Run()
}
