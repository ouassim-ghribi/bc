package main

import (
	"fmt"

	"github.com/ouassim-ghribi/bc/mrkl"
)

func main() {

	txs := mrkl.ReadJSON("tx.json")

	// _, txs := mrkl.GetMerkleRootAndTransactions()

	root := mrkl.BuildTree(txs)

	fmt.Println("Merkle Tree:")
	mrkl.PrintMerkleTree(root, "", false, true)

	proof := mrkl.GenerateMerkleProof(root, txs[2])
	// Verify the proof.
	if mrkl.VerifyMerkleProof(proof, txs[2], root.Hash) {
		fmt.Println("Merkle Proof is valid.")
	} else {
		fmt.Println("Merkle Proof is not valid.")
	}

}
