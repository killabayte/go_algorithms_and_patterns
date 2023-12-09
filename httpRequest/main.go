package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/go-resty/resty"
)

var (
	apiKey    = os.Getenv("BINANCE_API_KEY")
	secretKey = os.Getenv("BINANCE_SECRET_KEY")
	baseURL   = "https://api.binance.com"
)

func generateSignature(params map[string]string) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var queryString string
	for _, k := range keys {
		queryString += k + "=" + params[k] + "&"
	}
	queryString = strings.TrimSuffix(queryString, "&")

	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(queryString))
	signature := hex.EncodeToString(mac.Sum(nil))

	return signature
}

func makeBinanceRequest(client *resty.Client, endpoint string, params map[string]string) string {
	if endpoint != "/api/v3/exchangeInfo" {
		params["timestamp"] = fmt.Sprintf("%d", time.Now().Unix()*1000)
		params["signature"] = generateSignature(params)
	}

	response, err := client.R().
		SetQueryParams(params).
		SetHeader("X-MBX-APIKEY", apiKey).
		Get(baseURL + endpoint)

	if err != nil {
		return fmt.Sprintf("Error making request to %s: %s", endpoint, err)
	}

	return string(response.Body())
}

func main() {
	client := resty.New()

	// Request account information
	accountInfo := makeBinanceRequest(client, "/api/v3/account", map[string]string{})
	fmt.Println("Account Information:")
	fmt.Println(accountInfo)

	// Request spot transaction history for a specific symbol (replace "UNIUSDT" with the desired symbol)
	tradeHistory := makeBinanceRequest(client, "/api/v3/myTrades", map[string]string{"symbol": "UNIUSDT"})
	fmt.Println("\nSpot Trade History:")
	fmt.Println(tradeHistory)

	// Request exchange information
	exchangeInfo := makeBinanceRequest(client, "/api/v3/exchangeInfo", map[string]string{})
	fmt.Println("\nExchange Information:")
	fmt.Println(exchangeInfo)
}
