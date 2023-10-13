package main

import (
	"fmt"

	"github.com/ouassim-ghribi/bc/mrkl"
)

func main() {

	// txs := mrkl.ReadJSON("tx.json")

	_, txs := mrkl.GetMerkleRootAndTransactions()

	root := mrkl.BuildTreeH(txs)

	fmt.Println("Merkle Tree:")
	mrkl.PrintMerkleTree(root, "", false, true)

}
