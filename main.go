package main

import (
	"encoding/json"
	"log"

	"github.com/antunesgabriel/golang-blockchain/domain"
)

func main() {
	genesisData := map[string]string{
		"To":    "Genesis",
		"Value": "67.9",
	}

	genesis := domain.NewBlock(genesisData, "0")

	blockchain := domain.NewBlockChain(genesis)

	data := map[string]string{
		"From":    "Genesis",
		"To":      "Antunes",
		"Value":   "1.0",
		"Message": "Great Job!",
	}

	lastBlock := blockchain.GetLastBlock()

	block := domain.NewBlock(data, string(lastBlock.Hash))

	blockchain.AddBlock(block)

	j, _ := json.Marshal(block)

	log.Printf("Block Minerado %s", j)
}
