package handler

import (
	"blockchainETH-MongoDb/core/domain"
	"blockchainETH-MongoDb/modules"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"net/http"

	"github.com/gorilla/mux"
)

// ClientHandler ethereum client instance
type ClientHandler struct {
	*ethclient.Client
}

func (client ClientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	module := vars["module"]

	//address := r.URL.Query().Get("address")
	//hash := r.URL.Query().Get("hash")
	//id := r.URL.Query().Get("id")

	w.Header().Set("Content-Type", "application/json")

	switch module {
	case "latest-block":
		_block := modules.GetLatestBlock(*client.Client)
		json.NewEncoder(w).Encode(_block)

	case "get-tx":

		decoder := json.NewDecoder(r.Body)
		var h domain.HashRequest

		err := decoder.Decode(&h)

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&domain.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}

		if h.Hash == "" {
			json.NewEncoder(w).Encode(&domain.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}
		txHash := common.HexToHash(h.Hash)
		_tx := modules.GetTxByHash(*client.Client, txHash)

		if _tx != nil {
			json.NewEncoder(w).Encode(_tx)
			return
		}

		json.NewEncoder(w).Encode(&domain.Error{
			Code:    404,
			Message: "Tx Not Found!",
		})
	case "send-eth":
		decoder := json.NewDecoder(r.Body)
		var t domain.TransferEthRequest

		err := decoder.Decode(&t)

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&domain.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}
		_hash, err := modules.TransferEth(*client.Client, t.PrivKey, t.To, t.Amount)

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&domain.Error{
				Code:    500,
				Message: "Internal server error",
			})
			return
		}

		json.NewEncoder(w).Encode(&domain.HashResponse{
			Hash: _hash,
		})
	case "id":
		decoder := json.NewDecoder(r.Body)
		var b domain.BlockNumberRequest

		err := decoder.Decode(&b)

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&domain.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}
		fmt.Println("Block number: ", b.BlockNumber)

		_block := modules.GetBlockByNumber(*client.Client, b.BlockNumber)
		fmt.Println("Block: ", _block)
		json.NewEncoder(w).Encode(_block)

	case "get-balance":
		decoder := json.NewDecoder(r.Body)
		var b domain.BalanceRequest

		err := decoder.Decode(&b)

		if b.Address == "" {
			json.NewEncoder(w).Encode(&domain.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}

		balance, err := modules.GetAddressBalance(*client.Client, b.Address)

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&domain.Error{
				Code:    500,
				Message: "Internal server error",
			})
			return
		}

		json.NewEncoder(w).Encode(&domain.BalanceResponse{
			Address: b.Address,
			Balance: balance,
			Symbol:  "Ether",
			Units:   "Wei",
		})

	}

}
