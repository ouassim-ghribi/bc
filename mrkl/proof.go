package mrkl

import (
	"bytes"
	"fmt"
)

// GenerateMerkleProof generates a Merkle proof for a data element.
func GenerateMerkleProof(root *MerkleNode, data []byte) [][]byte {
	var proof [][]byte

	hash := CalculateHash(data)
	fmt.Printf("Data Hash: %x\n", hash)

	// Find the path to the data element by traversing the tree.
	currentNode := root
	for {
		proof = append(proof, currentNode.Hash)

		fmt.Printf("%x\n", currentNode.Hash)
		fmt.Println(bytes.Compare(hash, currentNode.Hash))

		if bytes.Compare(hash, currentNode.Hash) < 0 {
			if currentNode.Left == nil {
				break
			}
			currentNode = currentNode.Left
		} else {
			if currentNode.Right == nil {
				break
			}
			currentNode = currentNode.Right
		}
	}

	return proof
}

func VerifyMerkleProof(proof [][]byte, data []byte, rootHash []byte) bool {
	// Start with the data element's hash.
	hash := CalculateHash(data)

	// Iterate through the proof.
	for _, sibling := range proof {
		if bytes.Compare(data, sibling) < 0 {
			// Data is lexicographically smaller; concatenate data and sibling.
			hashData := append(data, sibling...)
			hash = CalculateHash(hashData)
		} else {
			// Sibling is lexicographically smaller; concatenate sibling and data.
			hashData := append(sibling, data...)
			hash = CalculateHash(hashData)
		}
		data = hash // Update data for the next iteration.
	}

	// Compare the final hash with the known root hash to verify the proof.
	return bytes.Compare(hash, rootHash) == 0
}
