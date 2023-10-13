package mrkl

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type BlockInfo struct {
	Hash      string        `json:"hash"`
	Ver       int64         `json:"ver"`
	PrevBlock string        `json:"prev_block"`
	MrklRoot  string        `json:"mrkl_root"`
	N_tx      int64         `json:"n_tx"`
	Tx        []Transaction `json:"tx"`
}

type Transaction struct {
	Hash string `json:"hash"`
	Ver  int    `json:"ver"`
}

func GetLatestBlock() string {
	resp, err := http.Get("https://blockchain.info/q/latesthash?cors=true")

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	blockHash := string(body)

	return blockHash
}

func GetMerkleRootAndTransactions() (string, []Transaction) {
	blockHash := GetLatestBlock()

	requestURL := fmt.Sprintf("https://blockchain.info/rawblock/%s?cors=true", blockHash)

	r, err := http.Get(requestURL)

	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var block BlockInfo
	if err := json.Unmarshal(body, &block); err != nil {
		log.Fatalln(err)
	}

	// Now you can access the data in the 'block' struct
	fmt.Printf("Block Hash: %s\n", block.Hash)
	fmt.Printf("Block Version: %d\n", block.Ver)
	fmt.Printf("Block Number of txs: %d\n", block.N_tx)
	fmt.Printf("Block First transaction: %v\n", block.Tx[0].Hash)

	return block.MrklRoot, block.Tx
}
