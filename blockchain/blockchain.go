package blockchain

type BlockChain struct {
	Blocks []*Block
}

func (chain *BlockChain) AddBlock(data string) {
	// get prev block
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	// create new
	new := CreateBlock(data, prevBlock.Hash)
	// add to chain
	chain.Blocks = append(chain.Blocks, new)
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
