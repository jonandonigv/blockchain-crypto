package blockchain

import (
	"log"

	"github.com/boltdb/bolt"
	"github.com/jonandonigv/blockchain-crypto/block"
)

const dbfile = "blockchain.db"
const blockBucket = "blocks"

type Blockchain struct {
	tip    []byte
	Blocks *bolt.DB
}

// Adds a new block into the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := block.NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// Creates the genesis block. The first block of a block-chain data structure
func NewGenesisBlock() *block.Block {
	return block.NewBlock("Genesis block", []byte{})
}

// Creates a new block-chain
func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbfile, 0600, nil)

	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic(err)
			}
			err = b.Put(genesis.Hash, genesis.Serialize())
			err = b.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}
		return nil
	})
	bc := Blockchain{tip, db}

	return &bc
	// return &Blockchain{[]*block.Block{NewGenesisBlock()}}
}
