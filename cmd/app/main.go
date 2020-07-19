package main

import (
	"github.com/gorilla/mux"
	"go-blockchain/internal"
	"go-blockchain/pkg/blockchain"
	"net/http"
)

func main() {
	bc := blockchain.New()
	bc.AddBlock("testing")

	bcHnd := internal.BlockchainHandler{
		bc,
	}

	r := mux.NewRouter()
	r.HandleFunc("/blocks", bcHnd.GetBlocks)
	r.HandleFunc("/mine", bcHnd.MineBlock)
	
	srv := http.Server{
		Handler: r,
		Addr: "0.0.0.0:8080",
	}

	_ = srv.ListenAndServe()
}
