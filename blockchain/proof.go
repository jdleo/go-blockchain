package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

// base difficulty
const Difficulty = 16

// proof of work instance
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// helper method to create a new proof of work target
func NewProof(b *Block) *ProofOfWork {
	// create big int, left shift
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	// pow instance
	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) InitData(nonce int) []byte {
	// create data from block
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	// loop from nonce 0 -> max
	for nonce < math.MaxInt64 {
		// create data to be hashed
		data := pow.InitData(nonce)
		// get hash
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		// check if we're under target (valid block)
		if intHash.Cmp(pow.Target) == -1 {
			// we found valid block
			break
		} else {
			// increment nonce
			nonce++
		}
	}

	fmt.Println()

	// return nonce and hash
	return nonce, hash[:]
}

// helper method to validate block
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	// create data
	data := pow.InitData(pow.Block.Nonce)

	// create hash, set bytes
	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	// check if under difficulty target
	return intHash.Cmp(pow.Target) == -1
}

// helper method to convert int64 into byte slice
func ToHex(num int64) []byte {
	buf := new(bytes.Buffer)

	// write to buffer
	err := binary.Write(buf, binary.BigEndian, num)

	if err != nil {
		log.Panic(err)
	}

	// return byte slice
	return buf.Bytes()
}
