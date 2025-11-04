package cli

import (
	"flag"
	"os"

	blockchain "github.com/jonandonigv/blockchain-crypto/block-chain"
)

type CLI struct {
	bc *blockchain.Blockchain
}

func (cli *CLI) printUsage() {}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) Run() {
	cli.validateArgs()

	addBlock := flag.NewFlagSet("addblock", flag.ExitOnError)
}
