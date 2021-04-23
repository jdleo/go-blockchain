package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

func CreateBlock(data string, prevHash []byte) *Block {
	// create new block
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	// create proof and mine
	pow := NewProof(block)
	nonce, hash := pow.Run()

	// set hash and nonce
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
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
