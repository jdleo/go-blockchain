package main

import (
	"os"

	wallet "github.com/jdleo/go-blockchain/wallet"
)

func main() {
	defer os.Exit(0)
	// cmd := cli.CommandLine{}
	// cmd.Run()

	w := wallet.MakeWallet()
	w.Address()
}
