package cmd

type Transactions struct {
	Total float64 `json:"btc_total"`
	Avg   float64 `json:"btc_avg"`
	Txs   []Tx    `json:"btc"`
}

type Tx [2]float64
