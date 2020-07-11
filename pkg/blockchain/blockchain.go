package blockchain

import (
	"bytes"
	"time"
)

type Blockchain struct {
	Blocks []*Block
}

func New() *Blockchain {
	blockchain := &Blockchain{
		Blocks: []*Block {Genesis()},
	}

	return blockchain
}

func (bc *Blockchain) AddBlock(data interface{}) {
	block := NewBlock(time.Now(), bc.Blocks[len(bc.Blocks)-1].Hash, data)
	bc.Blocks = append(bc.Blocks, block)
}

func (bc *Blockchain) ReplaceChain(chain []*Block) {
	if len(bc.Blocks) > len(chain) || !isValid(chain) {
		return
	}

	bc.Blocks = chain
}

func isValid(chain []*Block) bool {
	if chain[0] != Genesis() {
		return false
	}

	lastBlock := chain[0]
	for _, block := range chain[1:] {
		if !bytes.Equal(lastBlock.Hash, block.LastBlockHash) || !bytes.Equal(block.Hash, blockHash(block)) {
			return false
		}
	}

	return true
}