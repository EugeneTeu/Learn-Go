package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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

func errLogger(err error, message string) {
	if err != nil {
		log.Println(message)
		panic(err.Error())
	}
}

func main() {
	SymbolDetails := SymbolDetails{Symbols: []SymbolDetail{}}
	// tickerArray := GetTickers("https://api.gemini.com/v1/symbols")

	// ticker := make(map[string]map[string]interface{})

	// //fmt.Println(tickerArray)
	// for _, val := range tickerArray {

	// 	// strippedString := val[1 : len(val)-1]
	// 	fmt.Printf("Getting info for %s\n", val)
	// 	queryString := PrepareGetPubTicketString(val)
	// 	//fmt.Println(queryString)
	// 	tickerInfo := CallAPI(queryString)
	// 	ticker[val] = tickerInfo
	// }
	// for key, value := range ticker {
	// 	fmt.Printf("%s : %s\n", key, value)
	// }

	symbols := GetTickers("https://api.gemini.com/v1/symbols")

	for _, val := range symbols {
		fmt.Printf("Getting info for %s\n", val)
		queryString := PrepareGetPubTicketString(val)
		tickerInfo := GetSymbolDetail(queryString)
		SymbolDetails.Symbols = append(SymbolDetails.Symbols, tickerInfo)
	}

	priceFeed := GetPriceFeed()
	for _, value := range priceFeed {
		fmt.Printf("%+v\n", value)
	}

	for _, value := range SymbolDetails.Symbols {
		fmt.Printf("%+v\n", value)
	}

	//TODO: convert this to a api capable server
	// err := godotenv.Load()
	// errLogger(err, "error with loading env variables")
	// err = http.ListenAndServe(os.Getenv("API_PORT"), mux.NewRouter().StrictSlash(true))
	// errLogger(err, "could not start HTTP server")
	// log.Printf("HTTP server started at port: %s\n", os.Getenv("API_PORT"))
}

func PrepareGetPubTicketString(symbol string) (result string) {
	return "https://api.gemini.com/v1/symbols/details/" + symbol
}

func GetPriceFeedUrl() (result string) {
	return "https://api.gemini.com/v1/pricefeed"
}

func GetTickers(url string) (val []string) {
	// gemini returns [ "BTCUSD" ....]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	var response []string
	json.Unmarshal(result, &response)
	if err != nil {
		fmt.Println(err)
	}
	return response
}

func GetPriceFeed() (val []PriceFeedStruct) {
	// gemini returns [ "BTCUSD" ....]
	resp, err := http.Get(GetPriceFeedUrl())
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	var response []PriceFeedStruct
	json.Unmarshal(result, &response)
	if err != nil {
		fmt.Println(err)
	}
	return response
}

func GetSymbolDetail(url string) (val SymbolDetail) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	var response SymbolDetail
	json.Unmarshal(result, &response)
	if err != nil {
		fmt.Println(err)
	}
	return response
}

/* Reference */
func CallAPI(url string) (val map[string]interface{}) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	var response map[string]interface{}

	result, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(result, &response)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(response)
	return response
}
