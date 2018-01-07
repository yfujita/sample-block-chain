package block_chain

type BlockChain []*Block


func (chain *BlockChain)Add(block *Block) BlockChain {
	return append(*chain, block)
}

func (chain *BlockChain)Clone() BlockChain {
	newChain := make(BlockChain, len(*chain) + 1)
	for i, block := range *chain {
		newChain[i] = block.Clone()
	}
	return newChain
}

