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
}
