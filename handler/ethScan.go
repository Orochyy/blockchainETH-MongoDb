package handler

import (
	"blockchainETH-MongoDb/core/domain"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nanmu42/etherscan-api"
	"net/http"
)

type EthScanHandler struct {
	*etherscan.Client
}

func (e EthScanHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	module := vars["module"]

	w.Header().Set("Content-Type", "application/json")

	switch module {
	case "get-balance":
		balance, err := e.AccountBalance("0x281055afc982d96fab65b3a49cac8b878184cb16")
		if err != nil {
			panic(err)
		}
		// balance in wei, in *big.Int type
		fmt.Println(balance.Int())

		json.NewEncoder(w).Encode(balance)

	case "get-block":
		decoder := json.NewDecoder(r.Body)
		var b domain.BlockNumberRewardsRequest

		err := decoder.Decode(&b)

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&domain.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}

		block, err := e.BlockReward(b.BlockNumber)

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&domain.Error{
				Code:    500,
				Message: "Internal server error",
			})
			return
		}

		json.NewEncoder(w).Encode(block)

	case "transaction":
		decoder := json.NewDecoder(r.Body)

		var t domain.TransactionRequest

		err := decoder.Decode(&t)

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&domain.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}

		transaction, err := e.ERC1155Transfers(t.ContractAddress, t.Address, t.StartBlock, t.EndBlock, t.Page, t.Offset, t.Desc)

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&domain.Error{
				Code:    500,
				Message: "Internal server error",
			})
			return
		}
		json.NewEncoder(w).Encode(transaction)

	case "hash":
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
				Message: "Hash cannot be empty",
			})
			return
		}

		transaction := fmt.Sprintf("https://api.etherscan.io/api?module=proxy&action=eth_getTransactionByHash&txhash=%s&apikey=9PE85ZN1F87E9YMDP9NBF1MXWU4EA1QSYP", h.Hash)
		//https://api.etherscan.io/api?module=proxy&action=eth_getTransactionByHash&txhash=0x2b617a3097f078c6be5d39de3d8acfbf5b72bab4fad4da281af26302c67ea941&apikey=9PE85ZN1F87E9YMDP9NBF1MXWU4EA1QSYP

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&domain.Error{
				Code:    500,
				Message: "Internal server error",
			})
			return
		}

		json.NewEncoder(w).Encode(transaction)
		get, err := http.Get(transaction)
		if err != nil {
			return
		}
		fmt.Println(get)

	case "google":
		qq, err := http.Get("https://google.com")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(qq)
	}
}
