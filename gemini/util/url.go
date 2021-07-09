package util

func GetPubTicketString(symbol string) (result string) {
	return "https://api.gemini.com/v1/symbols/details/" + symbol
}

func GetPriceFeedUrl() (result string) {
	return "https://api.gemini.com/v1/pricefeed"
}

func GetWebSocketUrl(symbol string) (result string) {
	return "wss://api.gemini.com/v1/marketdata/" + symbol
}
