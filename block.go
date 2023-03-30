package HistChain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

type Block struct {
	Index        int64
	Timestamp    int64
	Event        Event
	PreviousHash string
	Hash         string
}

func (b *Block) setHash() {
	eventBytes, _ := json.Marshal(b.Event)
	hashData := string(b.Index) + string(b.Timestamp) + string(eventBytes) + b.PreviousHash
	hashInBytes := sha256.Sum256([]byte(hashData))
	b.Hash = hex.EncodeToString(hashInBytes[:])
}
