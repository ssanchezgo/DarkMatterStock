package model

type StockItem struct {
	Ticker     string `json:"ticker"`
	Company    string `json:"company"`
	Brokerage  string `json:"brokerage"`
	Action     string `json:"action"`
	TargetFrom string `json:"target_from"`
	TargetTo   string `json:"target_to"`
	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`
	Time       string `json:"time"`
}

type APIResponse struct {
	Items    []StockItem `json:"items"`
	NextPage string      `json:"next_page"`
}
