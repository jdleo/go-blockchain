package blockchain

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
