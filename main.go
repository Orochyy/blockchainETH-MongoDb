package main

import (
	"blockchainETH-MongoDb/handler"
	"github.com/nanmu42/etherscan-api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//client, err := ethclient.Dial("http://localhost:7545")
	//
	//if err != nil {
	//	fmt.Println(err)
	//}

	client := etherscan.New(etherscan.Mainnet, "9PE85ZN1F87E9YMDP9NBF1MXWU4EA1QSYP")

	r := mux.NewRouter()

	//r.Handle("/api/v1/eth/{module}", handler.ClientHandler{Client: client})

	r.Handle("/api/v1/eth/{module}", handler.EthScanHandler{Client: client})

	log.Fatal(http.ListenAndServe(":8080", r))
}
