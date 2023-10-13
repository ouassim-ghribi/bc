package main

import (
	"fmt"

	"github.com/strlrd-29/bc/mrkl"
)

func main() {

	txs := mrkl.ReadJSON("tx.json")

	// blockHash := mrkl.GetLatestBlock()

	// MrklRoot, Tx := mrkl.GetMerkleRootAndTransactions(blockHash)

	// fmt.Println(MrklRoot)
	// fmt.Println(Tx)

	root := mrkl.BuildTree(txs)

	fmt.Println("Merkle Tree:")
	mrkl.PrintMerkleTree(root, "", false, true)

}
