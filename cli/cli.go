package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	blockchain "github.com/jonandonigv/blockchain-crypto/block-chain"
)

type CLI struct {
	Bc *blockchain.Blockchain
}

func (cli *CLI) createBlockchain(address string) {
	bc := blockchain.CreateBlockchain(address)
	bc.Blocks.Close()
	fmt.Println("Done!")
}

func (cli *CLI) getBalance(address string) {
	bc := blockchain.NewBlockchain(address)
	defer bc.Blocks.Close()

	balance := 0
	UTXOs := bc.FindUTXO(address)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}

// I'm not sure when I stoped using this function
/* func (cli *CLI) addblock(data string) {
	cli.Bc.AddBlock(data)
	fmt.Println("Success!")
} */

func (cli *CLI) printUsage() {
	fmt.Println("Usage: ")
	fmt.Println(" getbalance -address ADDRESS - Get balance of ADDRESS")
	fmt.Println(" createblockchain -address ADDRESS - Create a blockchain and send genesis block reward to ADDRESS ")
	fmt.Println(" printchain -Print all the blocks of the blockchain")
	fmt.Println(" send -from FROM -to TO -amount AMOUNT - Send AMOUNT of coins from FROM to TO")
}

func (cli *CLI) printChain() {
	bci := cli.Bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) send(from, to string, amount int) {
	bc := blockchain.NewBlockchain(from)
	defer bc.Blocks.Close()

	tx := blockchain.NewUTXOTransaction(from, to, amount, bc)
	bc.AddBlock([]*blockchain.Transaction{tx})
	fmt.Println("Success!")
}

func (cli *CLI) Run() {
	cli.validateArgs()

	getBalancedCMD := flag.NewFlagSet("getbalance", flag.ExitOnError)
	createBlockchainedCMD := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	sendCMD := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	getBalancedAddress := getBalancedCMD.String("address", "", "The address to get balance for")
	createBlockchainAddress := createBlockchainedCMD.String("address", "", "The address to send genesis block reward to")
	sendFrom := sendCMD.String("from", "", "Source wallet address")
	sendTo := sendCMD.String("to", "", "Destination wallet address")
	sendAmount := sendCMD.Int("amount", 0, "Amount to send")

	// addBlock := flag.NewFlagSet("addblock", flag.ExitOnError)
	//
	// addBlockData := addBlock.String("data", "", "Block data")

	switch os.Args[1] {
	case "getbalance":
		err := getBalancedCMD.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createblockchain":
		err := createBlockchainedCMD.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "send":
		err := sendCMD.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if getBalancedCMD.Parsed() {
		if *getBalancedAddress == "" {
			getBalancedCMD.Usage()
			os.Exit(1)
		}
		cli.getBalance(*getBalancedAddress)
	}

	if createBlockchainedCMD.Parsed() {
		if *createBlockchainAddress == "" {
			createBlockchainedCMD.Usage()
			os.Exit(1)
		}
		cli.createBlockchain(*createBlockchainAddress)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

	if sendCMD.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
			sendCMD.Usage()
			os.Exit(1)
		}
		cli.send(*sendFrom, *sendTo, *sendAmount)
	}

}
