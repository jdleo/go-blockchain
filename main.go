package main

import (
	"os"

	cli "github.com/jdleo/go-blockchain/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}
