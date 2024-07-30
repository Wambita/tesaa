package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Transaction struct {
	From   string
	To     string
	Amount string
}

type Block struct {
	Hash     []byte
	Data     Transaction
	PrevHash []byte
	Nonce    int
}

func CreateBlock(data Transaction, prevHash []byte) *Block {
	block := &Block{[]byte{}, data, prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func Genesis() *Block {
	data := Transaction{
		"Tom",
		"Dan",
		"hello",
	}
	return CreateBlock(data, []byte{})
}

func (b *Block) Serialized() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	Handle(err)
	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	Handle(err)

	return &block
}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
