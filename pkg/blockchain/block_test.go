package blockchain

import (
	"encoding/hex"
	"testing"
	"time"
)

func TestNewBlock(t *testing.T) {
	timestamp := time.Unix(1594443757, 0)
	lastBlockHash := []byte("foo")
	data := "barz"

	block := NewBlock(timestamp, lastBlockHash, data)

	if hex.EncodeToString(block.Hash) != "5da43cf48f3bc6db7b7e5c3ca3161b2f63fc511a3c77fa0464db8b0464f89b5a"{
		t.Errorf("NewBlock(%d, %s, %s) failed. Expected hash: %s, got: %s",
			timestamp.Unix(), hex.EncodeToString(lastBlockHash), data,
			"5da43cf48f3bc6db7b7e5c3ca3161b2f63fc511a3c77fa0464db8b0464f89b5a", hex.EncodeToString(block.Hash))
	}
}