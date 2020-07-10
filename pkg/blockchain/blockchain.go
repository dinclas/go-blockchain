package blockchain

import "time"

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