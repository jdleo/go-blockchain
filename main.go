package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct {
	blocks []*Block
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
	prevBlock := chain.blocks[len(chain.blocks)-1]
	// create new
	new := CreateBlock(data, prevBlock.Hash)
	// add to chain
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First block after genesis")
	chain.AddBlock("Second block after genesis")
	chain.AddBlock("Third block after genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Block Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
