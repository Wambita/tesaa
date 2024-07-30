package main

import (
	"fmt"

	blockchain "blockchain/blockChain"
)

func main() {
	// Create a new blockchain instance with a mining difficulty of 2
	blockchain, err := blockchain.CreateBlockchain(2)
	if err != nil {
		fmt.Printf("Failed to create blockchain: %v\n", err)
		return
	}

	// Load existing blocks from the database
	err = blockchain.LoadBlocks()
	if err != nil {
		fmt.Printf("Failed to load blocks: %v\n", err)
		return
	}

	// Record transactions on the blockchain for Alice, Bob, and John
	blockchain.AddBlock("hey", "Bob", 7)
	blockchain.AddBlock("John", "Jane", 8)

	// Check if the blockchain is valid; expecting true
	fmt.Println("Blockchain valid:", blockchain.IsValid())
}
