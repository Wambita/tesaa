package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"

	"blockchain/blockChain"
)

type CommandLine struct {
	blockChain *blockchain.BlockChain
}

func (cli *CommandLine) printUsage() {
	fmt.Println("usage: ")
	fmt.Println(" add -block Block_Data - add a block to the chain")
	fmt.Println(" print - Prints the blocks in the chain")
}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) addBlock(data string){
	cli.blockChain.AddBlock(data)
	fmt.Println("Added Block!")
}

func (cli *CommandLine) printChain(){
	
}

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First Block After Genesis")
	chain.AddBlock("Second Block After Genesis")
	chain.AddBlock("Third Block After Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()

		pow := blockchain.NewProof(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
