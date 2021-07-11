package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sort"
	"strconv"
	"strings"
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
		processTopTenChangers(priceFeedArray)
	}()

	// wg.Add(1)
	// go func() {
	// 	openWSSConnection("ETHUSD")
	
	// 	wg.Done()
	// }()

	wg.Wait()

}

func processTopTenChangers(priceFeedArray *[]PriceFeedStruct) {
	if len(*priceFeedArray) < 10 {
		return
	}
	sort.Slice(*priceFeedArray, func(i, j int) bool {
		return (*priceFeedArray)[i].PercentChange24h < (*priceFeedArray)[j].PercentChange24h
	})

	topTenGreatestChange := (*priceFeedArray)[len(*priceFeedArray)-10:]
	result := []string{}
	for _, ticker := range topTenGreatestChange {
		val, _ := strconv.ParseFloat(ticker.PercentChange24h, 32)
		percentageChange := val * 100
		var percentageChangeStr string
		if percentageChange > 0 {
			percentageChangeStr = fmt.Sprintf("+%.2f", percentageChange)
		} else {
			percentageChangeStr = fmt.Sprintf("-%.2f", percentageChange)
		}
		str := fmt.Sprintf("%v, price: %v, 24h change: %v", ticker.Pair, ticker.Price, percentageChangeStr)
		result = append(result, str)
	}
	fmt.Println(result)
	log.Println(fmt.Sprintf("Top 10 movers today are:\n%s", strings.Join(result, "\n")))

}

func openWSSConnection(symbol string) bool {
	websocketConnection = ws.OpenWebSocket(symbol)
	defer websocketConnection.Close()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	done := make(chan struct{})

	go func() bool {
		defer close(done)
		for {
			_, _, err := websocketConnection.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return true
			}
			//log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return true
		case t := <-ticker.C:
			err := websocketConnection.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return true
			}
		case <-interrupt:
			log.Println("interrupt")

			err := websocketConnection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return true
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return true
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

	var arr []PriceFeedStruct

	for _, ticker := range response {
		if ticker.Pair[3:] == "USD" {
			arr = append(arr, ticker)
		}

	}

	return &arr
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
