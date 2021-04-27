package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	_ "github.com/gorilla/websocket"
)

type SymbolDetails struct {
	Symbols []SymbolDetail
	mutex   sync.Mutex
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
	SymbolDetails := SymbolDetails{Symbols: []SymbolDetail{}, mutex: sync.Mutex{}}
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
		go func(val string) {
			queryString := PrepareGetPubTicketString(val)
			tickerInfo := GetSymbolDetail(queryString)

			SymbolDetails.mutex.Lock()
			SymbolDetails.Symbols = append(SymbolDetails.Symbols, tickerInfo)
			SymbolDetails.mutex.Unlock()
		}(val)
	}

	priceFeed := GetPriceFeed()
	for _, value := range priceFeed {
		fmt.Printf("%+v\n", value)
	}

	// for _, value := range SymbolDetails.Symbols {
	// 	fmt.Printf("%+v\n", value)
	// }

	c := OpenWebSocket("ETHUSD")
	defer c.Close()
	_, message, err := c.ReadMessage()
	if err != nil {
		// handle error
	}
	//TODO: figure out how to decode message
	fmt.Println(message)

	//TODO: convert this to a api capable server
	// err := godotenv.Load()
	// errLogger(err, "error with loading env variables")
	// err = http.ListenAndServe(os.Getenv("API_PORT"), mux.NewRouter().StrictSlash(true))
	// errLogger(err, "could not start HTTP server")
	// log.Printf("HTTP server started at port: %s\n", os.Getenv("API_PORT"))
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

func OpenWebSocket(symbol string) (d *websocket.Conn) {
	c, _, err := websocket.DefaultDialer.Dial(GetWebSocketUrl(symbol), nil)
	if err != nil {
		errLogger(err, "something wrong w websocket")
	}
	return c
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

// websocket
//https://github.com/gorilla/websocket/tree/master/examples/echo
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	//... Use conn to send and receive messages.
}
