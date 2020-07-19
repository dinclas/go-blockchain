package internal

import (
	"encoding/json"
	"go-blockchain/pkg/blockchain"
	"net/http"
)

type BlockchainHandler struct {
	Bc *blockchain.Blockchain
}

func (hnd BlockchainHandler) GetBlocks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(hnd.Bc.Blocks)
	_, _ = w.Write(data)
}

func (hnd BlockchainHandler) MineBlock(w http.ResponseWriter, r *http.Request) {
	var data interface{}
	json.NewDecoder(r.Body).Decode(&data)
	hnd.Bc.AddBlock(data)
	hnd.GetBlocks(w, r)
}