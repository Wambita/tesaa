package main

import (
	"fmt"
	"blockchain/blockChain"
)

func main() {
	// Create a new blockchain instance with a mining difficulty of 2
	bc, err := blockChain.CreateBlockchain(2)
	if err != nil {
		fmt.Printf("Failed to create blockchain: %v\n", err)
		return
	}

	// Load existing blocks from the database
	err = bc.LoadBlocks()
	if err != nil {
		fmt.Printf("Failed to load blocks: %v\n", err)
		return
	}

	// Add blocks to the blockchain
	// err = bc.AddBlock(map[string]interface{}{"from": "Alice", "to": "Bob", "amount": 10})
	// if err != nil {
	// 	fmt.Printf("Failed to add block: %v\n", err)
	// 	return
	// }

	// err = bc.AddBlock(map[string]interface{}{"from": "Bob", "to": "Charlie", "amount": 5})
	// if err != nil {
	// 	fmt.Printf("Failed to add block: %v\n", err)
	// 	return
	// }

	// Validate the blockchain
	fmt.Println("Blockchain valid:", bc.IsValid())

	// View blocks in the blockchain
	fmt.Println("Viewing blocks:")
	err = bc.ViewBlocks()
	if err != nil {
		fmt.Printf("Failed to view blocks: %v\n", err)
		return
	}
}
