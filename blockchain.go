package main

import (
	"fmt"
	"sync"
	"time"
)

// Block defines the structure for each block in the blockchain
type Block struct {
	Timestamp     int64
	Data          []string
	PrevBlockHash string
	Hash          string
	Nonce         int
}

// Blockchain represents the structure for the blockchain, containing a slice of blocks
type Blockchain struct {
	blocks      []*Block
	difficulty  int
	mutex       sync.Mutex
	pendingData []string
}

// NewBlock creates and returns a new block given the data and the hash of the previous block
func NewBlock(data []string, prevBlockHash string, difficulty int) *Block {
	block := &Block{time.Now().Unix(), data, prevBlockHash, "", 0}
	pow := NewProofOfWork(block, difficulty)
	nonce, hash := pow.Run()

	block.Hash = fmt.Sprintf("%x", hash)
	block.Nonce = nonce

	return block
}

// AddBlock adds a new block to the blockchain with the given data
func (bc *Blockchain) AddBlock(data []string) {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()

	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash, bc.difficulty)

	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockchain creates a new blockchain with a genesis block
func NewBlockchain(difficulty int) *Blockchain {
	genesisBlock := NewBlock([]string{"Genesis Block"}, "", difficulty)

	return &Blockchain{
		blocks:     []*Block{genesisBlock},
		difficulty: difficulty,
	}
}

// Adding pedning data
func (bc *Blockchain) AddData(data string) {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()

	bc.pendingData = append(bc.pendingData, data)
}

// Mine a block with pending data
func (bc *Blockchain) MinePendingData() {
	bc.mutex.Lock()
	data := bc.pendingData
	bc.pendingData = []string{} // After mining, reinitialize data
	bc.mutex.Unlock()

	bc.AddBlock(data)
}

// GetLastBlock returns the last block in the blockchain
func (bc *Blockchain) GetLastBlock() *Block {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()

	return bc.blocks[len(bc.blocks)-1]
}
