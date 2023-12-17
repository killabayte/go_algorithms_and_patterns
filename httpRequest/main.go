package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/go-resty/resty"
)

type CommissionRates struct {
	Maker  string `json:"maker"`
	Taker  string `json:"taker"`
	Buyer  string `json:"buyer"`
	Seller string `json:"seller"`
}

type Balance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

type AccountInfo struct {
	MakerCommission            int             `json:"makerCommission"`
	TakerCommission            int             `json:"takerCommission"`
	BuyerCommission            int             `json:"buyerCommission"`
	SellerCommission           int             `json:"sellerCommission"`
	CommissionRates            CommissionRates `json:"commissionRates"`
	CanTrade                   bool            `json:"canTrade"`
	CanWithdraw                bool            `json:"canWithdraw"`
	CanDeposit                 bool            `json:"canDeposit"`
	Brokered                   bool            `json:"brokered"`
	RequireSelfTradePrevention bool            `json:"requireSelfTradePrevention"`
	PreventSor                 bool            `json:"preventSor"`
	UpdateTime                 int64           `json:"updateTime"`
	AccountType                string          `json:"accountType"`
	Balances                   []Balance       `json:"balances"`
	Permissions                []string        `json:"permissions"`
	UID                        int64           `json:"uid"`
}

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

func makeBinanceRequest(client *resty.Client, endpoint string, params map[string]string) []byte {
	params["timestamp"] = fmt.Sprintf("%d", time.Now().Unix()*1000)
	params["signature"] = generateSignature(params)

	response, err := client.R().
		SetQueryParams(params).
		SetHeader("X-MBX-APIKEY", apiKey).
		Get(baseURL + endpoint)

	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return response.Body()
}

func getTradeHistoryForSymbols(client *resty.Client, symbols []string) {
	var activeAssetsProcessed, inactiveAssetsProcessed int
	var wg sync.WaitGroup
	var mu sync.Mutex // Mutex to safely update counters

	for _, symbol := range symbols {
		wg.Add(1)
		go func(symbol string) {
			defer wg.Done()

			tradeHistory := makeBinanceRequest(client, "/api/v3/myTrades", map[string]string{"symbol": symbol + "USDT"})

			mu.Lock()
			defer mu.Unlock()

			// Update counters
			if len(tradeHistory) == 0 || string(tradeHistory) == "[]" || strings.Contains(string(tradeHistory), "Invalid symbol") {
				inactiveAssetsProcessed++
			} else {
				fmt.Printf("\nSpot Trade History for %s:\n", symbol)
				fmt.Println(string(tradeHistory))
				activeAssetsProcessed++
			}
		}(symbol)
		time.Sleep(150 * time.Millisecond)
	}

	wg.Wait()

	fmt.Println("\nTotal assets processed:", activeAssetsProcessed+inactiveAssetsProcessed)
	fmt.Println("Active assets processed:", activeAssetsProcessed)
	fmt.Println("Inactive assets processed:", inactiveAssetsProcessed)
}

func extractSymbols(accountInfo AccountInfo) []string {
	var symbols []string

	for _, balance := range accountInfo.Balances {
		// Assuming that non-zero balances indicate ownership of the asset
		if balance.Free != "0" || balance.Locked != "0" {
			symbols = append(symbols, balance.Asset)
		}
	}

	return symbols
}

func main() {
	client := resty.New()

	var accountInfo AccountInfo

	// Unmarshal the JSON data into the struct
	err := json.Unmarshal(makeBinanceRequest(client, "/api/v3/account", map[string]string{}), &accountInfo)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Access balances data
	for _, balance := range accountInfo.Balances {
		fmt.Printf("Asset: %s, Free: %s, Locked: %s\n", balance.Asset, balance.Free, balance.Locked)
	}

	// Extract symbols from accountInfo
	symbols := extractSymbols(accountInfo)
	fmt.Println("Symbols:", symbols)

	// Get trade history for the extracted symbols
	getTradeHistoryForSymbols(client, symbols)

}
