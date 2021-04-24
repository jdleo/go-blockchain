package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"

	ripemd160 "golang.org/x/crypto/ripemd160"
)

const (
	checksumLength = 4
	version        = byte(0x00)
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func (w *Wallet) Address() []byte {
	// get pub hash key
	pubHash := PublicKeyHash(w.PublicKey)

	// join version w/ pubhash
	versionedHash := append([]byte{version}, pubHash...)

	// create checksum
	checksum := Checksum(versionedHash)

	// join checksum with versioned hash
	fullHash := append(versionedHash, checksum...)

	// create base58 address
	address := Base58Encode(fullHash)

	fmt.Printf("Pub Key: %x\n", w.PublicKey)
	fmt.Printf("Pub Hash: %x\n", pubHash)
	fmt.Printf("Address: %s\n", address)

	return address
}

func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	// 256-bit elliptic curve output
	curve := elliptic.P256()
	// generate random private key (secure-random)
	private, err := ecdsa.GenerateKey(curve, rand.Reader)

	if err != nil {
		log.Panic(err)
	}

	// join x and y bytes together (this will be public key)
	pub := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pub
}

func MakeWallet() *Wallet {
	// get priv and pub keys
	private, public := NewKeyPair()

	// return new wallet reference
	return &Wallet{private, public}
}

func PublicKeyHash(pubKey []byte) []byte {
	// hash public key
	pubHash := sha256.Sum256(pubKey)

	// hash sha256 hash using ripemd
	hasher := ripemd160.New()

	_, err := hasher.Write(pubHash[:])
	if err != nil {
		log.Panic(err)
	}

	// add nothing else
	publicRipMD := hasher.Sum(nil)

	return publicRipMD
}

func Checksum(payload []byte) []byte {
	// hash payload twice
	firstHash := sha256.Sum256(payload)
	secondHash := sha256.Sum256(firstHash[:])

	// get first checksum length bytes
	return secondHash[:checksumLength]
}
