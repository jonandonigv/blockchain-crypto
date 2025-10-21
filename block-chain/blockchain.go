package blockchain

import "github.com/jonandonigv/blockchain-crypto/block"

type Blockchain struct {
	Blocks []*block.Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := block.NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// Creates the genesis block. The first block of a block-chain data structure
func NewGenesisBlock() *block.Block {
	return block.NewBlock("Genesis block", []byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*block.Block{NewGenesisBlock()}}
}
