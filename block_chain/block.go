package block_chain

import (
	"crypto/sha256"
	"errors"
	"strconv"
)

type Block struct {
	Nonce	int
	PreviousHash [32]byte
	Data 	string
	Hash    [32]byte
}

func (block *Block)Validate() bool {
	hash := CreateHash(block.Data, block.PreviousHash, block.Nonce)
	if HashString(hash) != HashString(block.Hash) {
		return false
	}
	return true
}

func (block *Block)Clone() *Block {
	newBlock := new(Block)
	newBlock.Nonce = block.Nonce
	newBlock.PreviousHash = block.PreviousHash
	newBlock.Data = block.Data
	newBlock.Hash = block.Hash
	return newBlock
}

func CreateBlock(data string, previousHash [32]byte, nonce int) (*Block, error) {
	hash := CreateHash(data, previousHash, nonce)
	if !ValidateHash(hash) {
		return nil, errors.New("Invalid hash value: " + HashString(hash))
	}

	block := new(Block)
	block.Nonce = nonce
	block.PreviousHash = previousHash
	block.Data = data
	block.Hash = hash
	return block, nil
}


func CreateHash(data string, previousHash [32]byte, nonce int) [32]byte {
	return sha256.Sum256([]byte(data + HashString(previousHash) + strconv.Itoa(nonce)))
}

func ValidateHash(hash [32]byte) bool {
	if hash[0] != 0 {
		return false
	} else if hash[1] != 0 {
		return false
	}
	return true
}

func HashString(hash [32]byte) string {
	buf := ""
	for _, v := range hash {
		buf += strconv.Itoa(int(v)) + ", "
	}
	return buf
}