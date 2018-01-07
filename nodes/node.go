package nodes

import "github.com/yfujita/sample-block-chain/block_chain"

type Node struct {
	chain block_chain.BlockChain
}

func (node *Node)ValidateChain(chain block_chain.BlockChain) bool {
	if len(chain) <= len(node.chain) {
		return false
	}
	var prevHash *[32]byte = nil
	for _, block := range chain {
		if !block.Validate() {
			return false
		}
		if prevHash != nil {
			if block_chain.HashString(block.PreviousHash) != block_chain.HashString(*prevHash) {
				return false
			}
		}
		prevHash = &block.Hash
	}
	return true
}

func (node *Node)UpdateChain(chain block_chain.BlockChain) {
	node.chain = chain
}
