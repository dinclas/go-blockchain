package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	Timestamp     time.Time
	Hash          []byte
	LastBlockHash []byte
	Data          interface{}
}

func NewBlock(timestamp time.Time, lastBlockHash []byte, data interface{}) *Block {
	return &Block{
		Timestamp:     timestamp,
		Hash:          hash(timestamp, lastBlockHash, data),
		LastBlockHash: lastBlockHash,
		Data:          data,
	}
}

func Genesis() *Block {
	return NewBlock(time.Now(), nil, nil)
}

func (b Block) String() string {
	jsonData, _ := json.Marshal(b.Data)
	return fmt.Sprintf("block={\n\t%d\n\t%s\n\t%s\n}", b.Timestamp.UnixNano(),
		jsonData,
		hex.EncodeToString(b.Hash))
}

func hash(timestamp time.Time, lastBlockHash []byte, data interface{}) []byte {
	msg, _ := json.Marshal([]interface{} {
		timestamp, lastBlockHash, data,
	})
	value := sha256.Sum256(msg)
	return value[:]
}

