package cli

import (
	"fmt"
	blockchain "github.com/jonandonigv/blockchain-crypto/block-chain"
	"strconv"
)

func (cli *CLI) printChain() {

	bc := blockchain.NewBlockchain("")
	defer bc.Blocks.Close()

	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Println("===================Block %x =====================\n", block.Hash)
		fmt.Println("Prev. block: %x\n", block.PrevBlockHash)
		pow := blockchain.NewProofOfWork(block)
		fmt.Println("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
		for _, tx := range block.Transaction {
			fmt.Println(tx)
		}
		fmt.Printf("\n\n")

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
