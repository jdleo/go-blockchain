package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

type Block struct {
	Hash         []byte
	Transactions []*Transaction
	PrevHash     []byte
	Nonce        int
}

// helper method to hash transactions in a block
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	// add each tx hash
	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}

	// join all byte slices and hash
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

func CreateBlock(txs []*Transaction, prevHash []byte) *Block {
	// create new block
	block := &Block{[]byte{}, txs, prevHash, 0}
	// create proof and mine
	pow := NewProof(block)
	nonce, hash := pow.Run()

	// set hash and nonce
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func Genesis(coinbase *Transaction) *Block {
	return CreateBlock([]*Transaction{coinbase}, []byte{})
}

// helper method to serialize block into byte slice for persistence
func (b *Block) Serialize() []byte {
	// set up buffer, and encoder
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	// encode our block struct
	err := encoder.Encode(b)

	Handle(err)

	return res.Bytes()
}

// helper method to deserialize block byte slice into struct
func (b *Block) Deserialize(data []byte) *Block {
	// set up block, and decoder
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))

	// decode and write to block struct
	err := decoder.Decode(&block)

	Handle(err)

	return &block
}

// helper method to handle error
func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
