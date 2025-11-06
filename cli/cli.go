package cli

import (
	"flag"
	"fmt"
	"log"
	"os"

	blockchain "github.com/jonandonigv/blockchain-crypto/block-chain"
)

type CLI struct {
	bc *blockchain.Blockchain
}

func (cli *CLI) addblock(data string) {
	cli.bc.AddBlock(data)
	fmt.Println("Success!")
}

func (cli *CLI) printUsage() {
	// TODO
}

func (cli *CLI) printChain() {
	// TODO
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) Run() {
	cli.validateArgs()

	addBlock := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := addBlock.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		err := addBlock.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlock.Parsed() {
		if *addBlockData == "" {
			addBlock.Usage()
			os.Exit(1)
		}
		cli.bc.AddBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}
