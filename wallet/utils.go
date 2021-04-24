package wallet

import (
	"log"

	"github.com/mr-tron/base58"
)

func Base58Encode(input []byte) []byte {
	encoded := base58.Encode(input)

	// return encoded base58 as byte slice
	return []byte(encoded)
}

func Base58Decode(input []byte) []byte {
	// attempt to decode input as string
	decoded, err := base58.Decode(string(input[:]))
	if err != nil {
		log.Panic(err)
	}

	return decoded
}
