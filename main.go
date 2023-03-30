package HistChain

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Initialize the blockchain with a genesis block
	blockchain := &Blockchain{
		Blocks: []*Block{
			createGenesisBlock(),
		},
	}

	// Add new blocks to the blockchain with events
	event1 := Event{Name: "Event 1", Date: "2023-01-01", Evidence: []string{"evidence1a", "evidence1b"}, UserFeedback: map[string]string{"user1": "feedback1"}}
	event2 := Event{Name: "Event 2", Date: "2023-02-15", Evidence: []string{"evidence2a", "evidence2b"}, UserFeedback: map[string]string{"user2": "feedback2"}}

	blockchain.addBlock(event1)
	blockchain.addBlock(event2)

	// Print the blockchain
	blockchainJSON, _ := json.MarshalIndent(blockchain.Blocks, "", "  ")
	fmt.Println("Blockchain:")
	fmt.Println(string(blockchainJSON))
}
