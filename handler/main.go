package handler

import (
	"blockchainETH-MongoDb/models"
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
	// Get parameter from url request
	vars := mux.Vars(r)
	module := vars["module"]
	id := vars["id"]

	// Get the query parameters from url request
	address := r.URL.Query().Get("address")
	hash := r.URL.Query().Get("hash")

	// Set our response header
	w.Header().Set("Content-Type", "application/json")

	// Handle each request using the module parameter:
	switch module {
	case "latest-block":
		_block := modules.GetLatestBlock(*client.Client)
		json.NewEncoder(w).Encode(_block)

	case "get-tx":
		if hash == "" {
			json.NewEncoder(w).Encode(&models.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}
		txHash := common.HexToHash(hash)
		_tx := modules.GetTxByHash(*client.Client, txHash)

		if _tx != nil {
			json.NewEncoder(w).Encode(_tx)
			return
		}

		json.NewEncoder(w).Encode(&models.Error{
			Code:    404,
			Message: "Tx Not Found!",
		})
	//case "id": for get block by id
	case "id":
		_block := modules.GetBlockByNumber(*client.Client, id)
		json.NewEncoder(w).Encode(_block)
		
	case "send-eth":
		decoder := json.NewDecoder(r.Body)
		var t models.TransferEthRequest

		err := decoder.Decode(&t)

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&models.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}
		_hash, err := modules.TransferEth(*client.Client, t.PrivKey, t.To, t.Amount)

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&models.Error{
				Code:    500,
				Message: "Internal server error",
			})
			return
		}

		json.NewEncoder(w).Encode(&models.HashResponse{
			Hash: _hash,
		})

	case "get-balance":
		if address == "" {
			json.NewEncoder(w).Encode(&models.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}

		balance, err := modules.GetAddressBalance(*client.Client, address)

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&models.Error{
				Code:    500,
				Message: "Internal server error",
			})
			return
		}

		json.NewEncoder(w).Encode(&models.BalanceResponse{
			Address: address,
			Balance: balance,
			Symbol:  "Ether",
			Units:   "Wei",
		})

	}

}