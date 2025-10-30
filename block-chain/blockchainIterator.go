package blockchain

import (
	"log"

	"github.com/boltdb/bolt"
	"github.com/jonandonigv/blockchain-crypto/block"
)

type BLockchainIterator struct {
	currentHash []byte
	Blocks      *bolt.DB
}

func (bc *Blockchain) Iterator() *BLockchainIterator {
	bci := &BLockchainIterator{bc.tip, bc.Blocks}
	return bci
}

func (i *BLockchainIterator) Next() *block.Block {
	var blck *block.Block

	err := i.Blocks.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		encodedBlock := b.Get(i.currentHash)
		blck = block.DeserializeBlock(encodedBlock)
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	i.currentHash = blck.PrevBlockHash

	return blck
}
