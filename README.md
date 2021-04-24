# Go Blockchain

Proof of concept Blockchain-based cryptocurrency written in Golang.

### Build CLI

```
git clone https://github.com/jdleo/go-blockchain.git
cd go-blockchain
make build
```

### Run

```
./main
```

### Usage

```
Usage:
 getbalance -address ADDRESS - get the balance for given address
 printchain - prints all blocks in the chain
 createblockchain -address ADDRESS - creates a blockchain locally, and genesis block is mined to address
 send -from FROM -to TO -amount AMOUNT - send AMOUNT from FROM to TO
 createwallet - Creates a new wallet
 listaddresses - lists all addresses stored locally in wallets.data
```
