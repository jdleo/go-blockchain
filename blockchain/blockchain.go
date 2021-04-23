package blockchain

import (
	"fmt"

	"github.com/dgraph-io/badger"
)

const (
	dbPath = "./tmp/blocks"
)

type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

type BlockChainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

func InitBlockChain() *BlockChain {
	var lastHash []byte

	// set up badgerDB
	opts := badger.DefaultOptions(dbPath)
	opts.Dir = dbPath
	opts.ValueDir = dbPath

	// attempt to open badgerDB
	db, err := badger.Open(opts)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		// check if we have last hash
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("No existing blockchain found")
			genesis := Genesis()
			fmt.Println("Genesis proved")

			// set genesis block, and LH in db
			err = txn.Set(genesis.Hash, genesis.Serialize())
			Handle(err)
			err = txn.Set([]byte("lh"), genesis.Hash)

			// get ref to last hash, break out
			lastHash = genesis.Hash
			return err
		} else {
			// get last hash
			item, err := txn.Get([]byte("lh"))
			Handle(err)
			lastHash, err = item.ValueCopy(lastHash)
			return err
		}
	})

	Handle(err)

	// initialize blockchain
	blockchain := BlockChain{lastHash, db}
	return &blockchain
}

func (chain *BlockChain) AddBlock(data string) {
	var lastHash []byte

	// get lasthash
	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		lastHash, err = item.ValueCopy(lastHash)

		return err
	})

	Handle(err)

	// create block based on last hash
	newBlock := CreateBlock(data, lastHash)

	// update block
	err = chain.Database.Update(func(txn *badger.Txn) error {
		// set block hash and lasthash
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		Handle(err)
		err = txn.Set([]byte("lh"), newBlock.Hash)

		chain.LastHash = newBlock.Hash

		return err
	})

	Handle(err)
}

func (chain *BlockChain) Iterator() *BlockChainIterator {
	iter := &BlockChainIterator{chain.LastHash, chain.Database}
	return iter
}

func (iter *BlockChainIterator) Next() *Block {
	var block *Block

	// open db
	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		Handle(err)
		encodedBlock, err := item.ValueCopy([]byte{})
		block = block.Deserialize(encodedBlock)

		return err
	})

	Handle(err)
	iter.CurrentHash = block.PrevHash
	return block
}
