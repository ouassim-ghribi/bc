package mrkl

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Hash  []byte
}

type Transactions struct {
	Tx []string `json:"tx"`
}

func NewNode(data []byte) *MerkleNode {
	return &MerkleNode{Left: nil, Right: nil, Hash: calculateHash(data)}
}

func calculateHash(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

// buildTree constructs the Merkle Tree from a list of data.
func BuildTree(dataList [][]byte) *MerkleNode {
	if len(dataList) == 0 {
		return nil
	}

	if len(dataList) == 1 {
		return NewNode(dataList[0])
	}

	var nodes []*MerkleNode

	// Create leaf nodes from the data.
	for _, data := range dataList {
		nodes = append(nodes, NewNode(data))
	}

	// Recursively build the tree by pairing and hashing nodes.
	for len(nodes) > 1 {
		var newNodes []*MerkleNode

		// Check if the length is odd and duplicate the last element.
		if len(nodes)%2 != 0 {
			nodes = append(nodes, nodes[len(nodes)-1])
		}

		for i := 0; i < len(nodes); i += 2 {
			left := nodes[i]
			right := nodes[i+1]

			hashData := append(left.Hash, right.Hash...)
			combinedHash := calculateHash(hashData)

			newNode := &MerkleNode{Left: left, Right: right, Hash: combinedHash}
			newNodes = append(newNodes, newNode)
		}
		nodes = newNodes
	}

	return nodes[0] // Return the root node.

}

func PrintMerkleTree(root *MerkleNode, prefix string, isLeft bool, isRoot bool) {
	if root != nil {
		fmt.Printf("%s", prefix)
		if isRoot {
			fmt.Print("└──")
			prefix += "    "
		} else {
			if isLeft {
				fmt.Print("├──")
				prefix += "│   "
			} else {
				fmt.Print("└──")
				prefix += "    "
			}
		}

		fmt.Printf("%x\n", root.Hash)
		PrintMerkleTree(root.Left, prefix, true, false)
		PrintMerkleTree(root.Right, prefix, false, false)
	}
}

func ReadJSON(path string) [][]byte {
	jsonFile, err := os.Open(path)

	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Close()

	bytes, err := io.ReadAll(jsonFile)

	if err != nil {
		log.Fatalln(err)
	}

	var tx Transactions
	if err := json.Unmarshal(bytes, &tx); err != nil {
		log.Fatalln(err)
	}

	var byteList [][]byte
	for _, str := range tx.Tx {
		byteList = append(byteList, []byte(str))
	}

	return byteList
}
