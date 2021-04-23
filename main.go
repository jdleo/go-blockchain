package main

import (
	"fmt"
	"strconv"

	blockchain "github.com/jdleo/go-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	for i := 0; i < 10; i++ {
		chain.AddBlock(fmt.Sprintf("Block #%d", i))
	}

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Block Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		// create new pow, and log validation
		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))

		fmt.Println()
	}
}
