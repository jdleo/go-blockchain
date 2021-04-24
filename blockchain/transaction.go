package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

type TxOutput struct {
	Value  int
	PubKey string
}

type TxInput struct {
	ID  []byte
	Out int
	Sig string
}

func (tx *Transaction) SetID() {
	// set up buffer
	var encoded bytes.Buffer
	var hash [32]byte

	// encode transaction data
	encode := gob.NewEncoder(&encoded)
	err := encode.Encode(tx)
	Handle(err)

	// set id to be hash of tx data
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

func CoinbaseTx(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", to)
	}

	// create input and output
	txin := TxInput{[]byte{}, -1, data}
	txout := TxOutput{100, to}

	// create txn
	tx := Transaction{nil, []TxInput{txin}, []TxOutput{txout}}
	// set id
	tx.SetID()

	return &tx
}

func (tx *Transaction) IsCoinbase() bool {
	// verify if this transaction is a coinbase transaction if
	// 1. only one input
	// 2. that input's ID is nil
	// 3. that input's out is -1 (flag for coinbase tx)
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].Out == -1
}

func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
