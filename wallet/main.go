package main

import (
	"bufio"
	"os"
	"wallet/core"
)

func main() {
	wallet := core.NewWallet()
	scanner := bufio.NewScanner(os.Stdin)
	runner := core.NewWalletApplicationRunner(wallet, scanner)
	runner.Run()
}
