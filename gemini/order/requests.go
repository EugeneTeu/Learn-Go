package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var (
	client *http.Client
)

type PayloadPrivateAPI struct {
	Request string `json:"request"`
	Nonce   string `json:"nonce"`
	OrderID int64  `json:"order_id,omitempty"`
}

func main() {
	client := &http.Client{}
	err := godotenv.Load()
	if err != nil {
		log.Println(err.Error())
		return
	}
	url := "https://api.gemini.com/"
	request := "/v1/mytrades"

	fullURL := url + request

	apiSecret := os.Getenv("API_SECRET")
	apiKey := os.Getenv("API_KEY")
	timestamp := time.Now()
	nonce := string(rune(timestamp.Local().Minute() * 1000))
	payload := &PayloadPrivateAPI{
		Request: request,
		Nonce:   nonce,
	}

	encodedPayload, err := json.Marshal(payload)
	if err != nil {
		log.Println(err.Error())
		return
	}
	b64 := b64.StdEncoding.EncodeToString(encodedPayload)

	signatureHash := hmac.New(sha512.New384, []byte(apiSecret))
	signatureHash.Write([]byte(b64))
	signature := hex.EncodeToString(signatureHash.Sum(nil))

	buf := bytes.NewBuffer([]byte(encodedPayload))
	req, _ := http.NewRequest("POST", fullURL, buf)
	req.Header.Set("X-GEMINI-APIKEY", apiKey)
	req.Header.Set("X-GEMINI-PAYLOAD", b64)
	req.Header.Set("X-GEMINI-SIGNATURE", signature)
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Content-length", "0")
	res, _ := client.Do(req)
	defer res.Body.Close()

	res_body, _ := ioutil.ReadAll(res.Body)
	log.Printf("%s", string(res_body))
}
