package main

type SymbolDetails struct {
	Symbols []SymbolDetail
}

type SymbolDetail struct {
	Symbol         string `json:"symbol"`
	BaseCurrency   string `json:"base_currency"`
	QuoteCurrency  string `json:"quote_currency"`
	TickSize       int    `json:"tick_size"`
	QuoteIncrement int    `json:"quote_increment"`
	MinOrderSize   string `json:"min_order_size"`
	Status         string `json:"status"`
}

type PriceFeedStruct struct {
	Pair             string `json:"pair"`
	Price            string `json:"price"`
	PercentChange24h string `json:"percentChange24h"`
}