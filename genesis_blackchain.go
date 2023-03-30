package HistChain

import "time"

func createGenesisBlock() *Block {
	genesisBlock := &Block{
		Index:        0,
		Timestamp:    time.Now().Unix(),
		Event:        Event{Name: "Genesis Event", Date: "", Evidence: []string{}, UserFeedback: map[string]string{}},
		PreviousHash: "0",
	}
	genesisBlock.setHash()
	return genesisBlock
}
