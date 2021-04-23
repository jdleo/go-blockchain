package blockchain

type BlockChain struct {
	Blocks []*Block
}

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
