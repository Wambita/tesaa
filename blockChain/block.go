package blockChain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Hash         string
	PreviousHash string
	Data         map[string]interface{}
	Timestamp    time.Time
	Pow          int
}

type Blockchain struct {
	Chain      []Block
	Difficulty int
	FilePath   string
}

// InitializeBlockchainFile sets up the JSON file for the blockchain
func InitializeBlockchainFile(filePath string) error {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		// File does not exist, create it and write an empty array
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		_, err = file.Write([]byte("[]"))
		if err != nil {
			return err
		}
		file.Close()
	}
	return nil
}

// NewBlock creates a new block with the given data and previous hash
func NewBlock(data map[string]interface{}, previousHash string, difficulty int) Block {
	block := Block{
		Data:         data,
		PreviousHash: previousHash,
		Timestamp:    time.Now(),
	}
	block.Pow = block.mine(difficulty)
	block.Hash = block.calculateHash()
	return block
}

// calculateHash computes the hash of the block
func (b *Block) calculateHash() string {
	data, _ := json.Marshal(b.Data)
	blockData := b.PreviousHash + string(data) + b.Timestamp.String() + strconv.Itoa(b.Pow)
	hash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", hash)
}

// mine performs proof-of-work for the block
func (b *Block) mine(difficulty int) int {
	prefix := strings.Repeat("0", difficulty)
	for !strings.HasPrefix(b.calculateHash(), prefix) {
		b.Pow++
	}
	return b.Pow
}

// CreateBlockchain initializes a new blockchain with a genesis block or loads from existing file
func CreateBlockchain(difficulty int, filePath string) (*Blockchain, error) {
	err := InitializeBlockchainFile(filePath)
	if err != nil {
		return nil, err
	}

	blockchain := &Blockchain{
		Difficulty: difficulty,
		FilePath:   filePath,
	}

	// Load existing blocks from file
	err = blockchain.LoadBlocks()
	if err != nil {
		if os.IsNotExist(err) {
			// File does not exist; create genesis block
			genesisBlock := Block{
				Hash:         "0",
				PreviousHash: "",
				Data:         map[string]interface{}{},
				Timestamp:    time.Now(),
				Pow:          0,
			}
			genesisBlock.Hash = genesisBlock.calculateHash()

			blockchain.Chain = append(blockchain.Chain, genesisBlock)

			// Store the genesis block
			err = blockchain.storeBlocks()
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return blockchain, nil
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data map[string]interface{}) error {
	lastBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := NewBlock(data, lastBlock.Hash, bc.Difficulty)
	bc.Chain = append(bc.Chain, newBlock)

	return bc.storeBlocks()
}

// storeBlocks saves all blocks to the JSON file
func (bc *Blockchain) storeBlocks() error {
	file, err := os.Create(bc.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(bc.Chain, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	return err
}

// LoadBlocks loads blocks from the JSON file into the blockchain
func (bc *Blockchain) LoadBlocks() error {
	file, err := os.Open(bc.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	var chain []Block
	err = json.Unmarshal(data, &chain)
	if err != nil {
		return err
	}

	bc.Chain = chain
	return nil
}

// IsValid checks if the blockchain is valid
func (bc *Blockchain) IsValid() bool {
	if len(bc.Chain) == 0 {
		return false
	}

	for i := 1; i < len(bc.Chain); i++ {
		currentBlock := bc.Chain[i]
		previousBlock := bc.Chain[i-1]

		if currentBlock.Hash != currentBlock.calculateHash() {
			return false
		}

		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}

	return true
}
