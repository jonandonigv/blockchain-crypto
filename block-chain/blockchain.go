package blockchain

import (
	"log"
	"os"

	"github.com/boltdb/bolt"
	"github.com/jonandonigv/blockchain-crypto/block"
	"github.com/jonandonigv/blockchain-crypto/transactions"
)

const dbfile = "blockchain.db"
const blockBucket = "blocks"
const genesisCoinbaseData = "The times 03/Jan/2009 Chancellor on brink of second bailout for banks"

type Blockchain struct {
	tip    []byte
	Blocks *bolt.DB
}

// Adds a new block into the blockchain
func (bc *Blockchain) AddBlock(data string) {

	var lastHash []byte

	err := bc.Blocks.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		lastHash = b.Get([]byte("l"))
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	newBlock := block.NewBlock(data, lastHash)

	err = bc.Blocks.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}
		err = b.Put([]byte("l"), newBlock.Hash)
		bc.tip = newBlock.Hash
		return nil
	})

}

// Creates the genesis block. The first block of a block-chain data structure
func NewGenesisBlock(coinbase *transactions.Transaction) *block.Block {
	return block.NewBlock([]*transactions.Transaction{coinbase}, []byte{})
}

func dbExist() bool {
	if _, err := os.Stat(dbfile); os.IsNotExist(err) {
		return false
	}
	return true
}

// Creates a new block-chain
func NewBlockchain(address string) *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbfile, 0600, nil)

	if err != nil {
		log.Panic(err)
	}
	// TODO: Update the createion of a new blockchain so it uses transactions
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
