package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	util "gemini/util"

	ws "gemini/websocket"

	"github.com/gorilla/websocket"
)

var (
	priceFeedArray = &[]PriceFeedStruct{}

	websocketConnection *websocket.Conn
)

func main() {
	SymbolDetails := SymbolDetails{Symbols: []SymbolDetail{}}
	symbols := GetTickers("https://api.gemini.com/v1/symbols")
	var wg sync.WaitGroup
	wg.Add(len(symbols))
	log.Println("Getting symbols")
	for _, val := range symbols {
		go func(val string) {
			defer wg.Done()
			queryString := util.GetPubTicketString(val)
			tickerInfo := GetSymbolDetail(queryString)
			SymbolDetails.Symbols = append(SymbolDetails.Symbols, tickerInfo)
		}(val)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		priceFeedArray = GetPriceFeed()
	}()
	wg.Wait()
	log.Println(fmt.Sprintf("Queried a total of %v symbols", len(SymbolDetails.Symbols)))
	log.Println(priceFeedArray)
	websocketConnection = ws.OpenWebSocket("ETHUSD")
	wg.Add(1)
	defer websocketConnection.Close()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := websocketConnection.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := websocketConnection.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")
			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := websocketConnection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}

}

func ErrLogger(err error, message string) {
	if err != nil {
		log.Println(message)
		panic(err.Error())
	}
}

func getData(val string) SymbolDetail {
	fmt.Printf("Getting info for %s\n", val)
	queryString := util.GetPubTicketString(val)
	tickerInfo := GetSymbolDetail(queryString)
	return tickerInfo
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

func GetPriceFeed() (val *[]PriceFeedStruct) {
	// gemini returns [ "BTCUSD" ....]
	resp, err := http.Get(util.GetPriceFeedUrl())
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
	return &response
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
