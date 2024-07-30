package main

import (
	"fmt"
	"blockchain/blockChain"
)

func main() {
	// Define the path to the JSON file
	filePath := "./blockchain.json"

	// Create a new blockchain instance with a mining difficulty of 3
	bc, err := blockChain.CreateBlockchain(3, filePath)
	if err != nil {
		fmt.Printf("Failed to create blockchain: %v\n", err)
		return
	}

	// Add a block to the blockchain
	err = bc.AddBlock(map[string]interface{}{"for": "Alice", "to": "Bob", "amount": 10, "name": "Tomlee"})
	if err != nil {
		fmt.Printf("Failed to add block: %v\n", err)
		return
	}

	// Validate the blockchain
	fmt.Println("Blockchain valid:", bc.IsValid())

	// View blocks in the blockchain
	fmt.Println("Viewing blocks:")
	for _, block := range bc.Chain {
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Data: %v\n", block.Data)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Proof of Work: %d\n", block.Pow)
		fmt.Println()
	}
}
