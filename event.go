package HistChain

import (
	"fmt"
	"time"
)

type Event struct {
	EventID      string            // Unique ID of the historical event
	Name         string            // Name of the historical event
	Date         string            // Date of the historical event (e.g. "2023-01-01")
	Evidence     []string          // List of evidence supporting the event
	UserFeedback map[string]string // User feedback on the event, with the key being the user ID and the value being the feedback
}

func (bc *Blockchain) addEvent(event Event) {
	if !isValidEvent(event) {
		fmt.Println("Event validation failed:", event.Name)
		return
	}

	newBlock := &Block{
		Index:        int64(len(bc.Blocks)),
		Timestamp:    time.Now().Unix(),
		Event:        event,
		PreviousHash: bc.Blocks[len(bc.Blocks)-1].Hash,
	}
	newBlock.setHash()
	bc.Blocks = append(bc.Blocks, newBlock)
}

func isValidEvent(event Event) bool {
	// Placeholder for validation logic
	// Implement consensus algorithm, social media integration, user feedback, and reputation systems here
	isValid := calculateConsensus(event.EventID)

	// If the event is valid, distribute tokens to users who contributed to the validation process
	if isValid {
		// Distribute tokens to users based on their reputation and the quality of their contributions
	}

	return isValid
}
