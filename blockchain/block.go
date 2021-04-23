package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	// join slices of bytes from previous block
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})

	// hash it
	hash := sha256.Sum256(info)

	// load hash
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	// create new block
	block := &Block{[]byte{}, []byte(data), prevHash}

	// derive hash from prev block
	block.DeriveHash()

	return block
}

func (chain *BlockChain) AddBlock(data string) {
	// get prev block
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	// create new
	new := CreateBlock(data, prevBlock.Hash)
	// add to chain
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
