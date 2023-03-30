package HistChain

import "time"

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) addBlock(event Event) {
	newBlock := &Block{
		Index:        int64(len(bc.Blocks)),
		Timestamp:    time.Now().Unix(),
		Event:        event,
		PreviousHash: bc.Blocks[len(bc.Blocks)-1].Hash,
	}
	newBlock.setHash()
	bc.Blocks = append(bc.Blocks, newBlock)
}
