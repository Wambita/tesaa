package blockChain

import (
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
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
	DB         *sql.DB
}

// InitializeDatabase sets up the SQLite database and schema
func InitializeDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./blockchain.db")
	if err != nil {
		return nil, err
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS blocks (
		hash TEXT PRIMARY KEY,
		previous_hash TEXT,
		data TEXT,
		timestamp TEXT,
		pow INTEGER
	);
	`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}
	return db, nil
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

// CreateBlockchain initializes a new blockchain with a genesis block
func CreateBlockchain(difficulty int) (*Blockchain, error) {
	db, err := InitializeDatabase()
	if err != nil {
		return nil, err
	}

	blockchain := &Blockchain{
		Difficulty: difficulty,
		DB:         db,
	}

	genesisBlock := Block{
		Hash:         "0",
		PreviousHash: "",
		Data:         map[string]interface{}{},
		Timestamp:    time.Now(),
		Pow:          0,
	}
	genesisBlock.Hash = genesisBlock.calculateHash()

	blockchain.Chain = append(blockchain.Chain, genesisBlock)

	err = blockchain.storeBlock(genesisBlock)
	if err != nil {
		return nil, err
	}

	return blockchain, nil
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data map[string]interface{}) error {
	lastBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := NewBlock(data, lastBlock.Hash, bc.Difficulty)
	bc.Chain = append(bc.Chain, newBlock)

	return bc.storeBlock(newBlock)
}

// storeBlock saves a block to the database
func (bc *Blockchain) storeBlock(block Block) error {
	data, _ := json.Marshal(block.Data)
	_, err := bc.DB.Exec(
		`INSERT INTO blocks (hash, previous_hash, data, timestamp, pow) VALUES (?, ?, ?, ?, ?)`,
		block.Hash,
		block.PreviousHash,
		string(data),
		block.Timestamp.Format(time.RFC3339),
		block.Pow,
	)
	return err
}


func (bc *Blockchain) ViewBlocks() error {
	rows, err := bc.DB.Query(`SELECT hash, previous_hash, data, timestamp, pow FROM blocks`)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var block Block
		var data string
		var timestamp string

		err := rows.Scan(&block.Hash, &block.PreviousHash, &data, &timestamp, &block.Pow)
		if err != nil {
			return err
		}

		block.Data = make(map[string]interface{})
		err = json.Unmarshal([]byte(data), &block.Data)
		if err != nil {
			return err
		}
		block.Timestamp, _ = time.Parse(time.RFC3339, timestamp)

		// Print block details
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Data: %v\n", block.Data)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Proof of Work: %d\n", block.Pow)
		fmt.Println()
	}

	return nil
}


// LoadBlocks loads blocks from the database into the blockchain
func (bc *Blockchain) LoadBlocks() error {
	rows, err := bc.DB.Query(`SELECT hash, previous_hash, data, timestamp, pow FROM blocks`)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var block Block
		var data string
		var timestamp string

		err := rows.Scan(&block.Hash, &block.PreviousHash, &data, &timestamp, &block.Pow)
		if err != nil {
			return err
		}

		block.Data = make(map[string]interface{})
		json.Unmarshal([]byte(data), &block.Data)
		block.Timestamp, _ = time.Parse(time.RFC3339, timestamp)

		bc.Chain = append(bc.Chain, block)
	}
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
