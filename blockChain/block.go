package blockchain

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
	data         map[string]interface{}
	hash         string
	previousHash string
	timestamp    time.Time
	pow          int
}

type Blockchain struct {
	genesisBlock Block
	chain        []Block
	difficulty   int
	db           *sql.DB
}

// Create a new SQLite database and initialize the schema
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

func (b *Block) calculateHash() string {
	data, _ := json.Marshal(b.data)
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.calculateHash()
	}
}

func CreateBlockchain(difficulty int) (*Blockchain, error) {
	db, err := InitializeDatabase()
	if err != nil {
		return nil, err
	}

	genesisBlock := Block{
		hash:      "0",
		timestamp: time.Now(),
	}

	blockchain := &Blockchain{
		genesisBlock: genesisBlock,
		chain:        []Block{genesisBlock},
		difficulty:   difficulty,
		db:           db,
	}

	// Store the genesis block in the database
	err = blockchain.storeBlock(genesisBlock)
	if err != nil {
		return nil, err
	}

	return blockchain, nil
}

func (b *Blockchain) storeBlock(block Block) error {
	data, _ := json.Marshal(block.data)
	_, err := b.db.Exec(
		`INSERT INTO blocks (hash, previous_hash, data, timestamp, pow) VALUES (?, ?, ?, ?, ?)`,
		block.hash,
		block.previousHash,
		string(data),
		block.timestamp.Format(time.RFC3339),
		block.pow,
	)
	return err
}

func (b *Blockchain) LoadBlocks() error {
	rows, err := b.db.Query(`SELECT hash, previous_hash, data, timestamp, pow FROM blocks`)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var block Block
		var data string
		var timestamp string

		err := rows.Scan(&block.hash, &block.previousHash, &data, &timestamp, &block.pow)
		if err != nil {
			return err
		}

		block.data = make(map[string]interface{})
		json.Unmarshal([]byte(data), &block.data)
		block.timestamp, _ = time.Parse(time.RFC3339, timestamp)

		b.chain = append(b.chain, block)
	}
	return nil
}

func (b *Blockchain) AddBlock(from, to string, amount float64) {
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}

	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
		data:         blockData,
		previousHash: lastBlock.hash,
		timestamp:    time.Now(),
	}

	newBlock.mine(b.difficulty)
	b.chain = append(b.chain, newBlock)

	// Store the new block in the database
	err := b.storeBlock(newBlock)
	if err != nil {
		fmt.Printf("Failed to store block: %v\n", err)
	}
}

func (b *Blockchain) IsValid() bool {
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]
		if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}
