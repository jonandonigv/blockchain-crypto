package blockchain

import (
	"fmt"
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

func (bc *Blockchain) FindUnspentTransactions(address string) []transactions.Transaction {
	var unspentTXs []transactions.Transaction
	// TODO: Add logic here
	return unspentTXs
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

	// TODO: Change data to transactions
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

// NewBlockchain creates a new Blockchain with a genesis block
func NewBlockchain(address string) *Blockchain {

	if dbExist() == false {
		fmt.Println("No existing blockchain found. Create one first.")
		os.Exit(1)
	}

	var tip []byte
	db, err := bolt.Open(dbfile, 0600, nil)

	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		tip = b.Get([]byte("l"))

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	bc := Blockchain{tip, db}

	return &bc
}

// CreateBlockchain creates a new blockchain db
func CreateBlockchain(address string) *Blockchain {
	if dbExist() {
		fmt.Println("Blockchain already exist.")
		os.Exit(1)
	}

	var tip []byte
	db, err := bolt.Open(dbfile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		cbtx := transactions.NewCoinbaseTX(address, genesisCoinbaseData)
		genesis := NewGenesisBlock(cbtx)

		b, err := tx.CreateBucket([]byte(blockBucket))
		if err != nil {
			log.Panic(err)
		}

		err = b.Put(genesis.Hash, genesis.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), genesis.Hash)
		if err != nil {
			log.Panic(err)
		}

		tip = genesis.Hash

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	bc := Blockchain{tip, db}

	return &bc
}
