package main

import (
	"crypto/sha256"
	"math/big"
	"strconv"
	"time"
)

const maxNonce = 1000000

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// NewProofOfWork creates a now proof of work for a block.
func NewProofOfWork(block *Block, difficulty int) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty))
	return &ProofOfWork{block, target}
}

// Run does the mining by finding a suitable nounce.
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	for nonce < maxNonce {
		select {
		case <-time.After(time.Second):
			// CPU security
			return 0, nil
		default:
			nonce++
			data := pow.prepareData(nonce)
			hash = sha256.Sum256(data)
			hashInt.SetBytes(hash[:])

			if hashInt.Cmp(pow.Target) == -1 {
				return nonce, hash[:]
			}
		}
	}

	return 0, nil
}

// Verify if the nounce is valid regarding to the proof of work
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := pow.prepareData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	return hashInt.Cmp(pow.Target) == -1
}

// prepareData prepares the data to be hashed, including the hash of the previous block.
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := append([]byte{}, pow.Block.PrevBlockHash...) // Create a new slice from PrevBlockHash
	for _, message := range pow.Block.Data {
		data = append(data, []byte(message)...)
	}
	data = append(data, IntToHex(int64(nonce))...)
	return data
}

func IntToHex(n int64) []byte {
	return []byte(strconv.FormatInt(n, 16))
}
