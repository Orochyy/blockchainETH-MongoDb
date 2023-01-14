package domain

// Block data structure
type Block struct {
	BlockNumber       int64         `json:"blockNumber"`
	Timestamp         uint64        `json:"timestamp"`
	Difficulty        uint64        `json:"difficulty"`
	Hash              string        `json:"hash"`
	TransactionsCount int           `json:"transactionsCount"`
	Transactions      []Transaction `json:"transactions"`
}

// Transaction data structure
type Transaction struct {
	Hash     string `json:"hash"`
	Value    string `json:"value"`
	Gas      uint64 `json:"gas"`
	GasPrice uint64 `json:"gasPrice"`
	Nonce    uint64 `json:"nonce"`
	To       string `json:"to"`
	Pending  bool   `json:"pending"`
}

// TransferEthRequest data structure
type TransferEthRequest struct {
	PrivKey string `json:"privKey"`
	To      string `json:"to"`
	Amount  int64  `json:"amount"`
}

// HashResponse data structure
type HashResponse struct {
	Hash string `json:"hash"`
}

type HashRequest struct {
	Hash string `json:"hash"`
}

// BalanceRequest data structure
type BalanceRequest struct {
	Address string `json:"address"`
}

// BalanceResponse data structure
type BalanceResponse struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
	Symbol  string `json:"symbol"`
	Units   string `json:"units"`
}

// Error data structure
type Error struct {
	Code    uint64 `json:"code"`
	Message string `json:"message"`
}

// BlockNumberRequest data structure
type BlockNumberRequest struct {
	BlockNumber int64 `json:"blockNumber"`
}

type BlockNumberRewardsRequest struct {
	BlockNumber int `json:"blockNumber"`
}

type BlockNumberRewardsResponse struct {
	BlockNumber int64  `json:"blockNumber"`
	TimeStamp   uint64 `json:"timeStamp"`
	BlockMiner  string `json:"blockMiner"`
	BlockReward string `json:"blockReward"`
}

type TransactionRequest struct {
	ContractAddress *string `json:"contractAddress"`
	Address         *string `json:"address"`
	StartBlock      *int    `json:"startBlock"`
	EndBlock        *int    `json:"endBlock"`
	Page            int     `json:"page"`
	Offset          int     `json:"offset"`
	Desc            bool    `json:"desc"`
}
