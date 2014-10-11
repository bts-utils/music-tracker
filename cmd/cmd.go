package cmd

type Transactions struct {
	Total int64 `json:"btc_total"`
	Avg   int64 `json:"btc_avg"`
	Txs   []Tx  `json:"btc"`
}

type Tx [2]float64
