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
	apiKey    = os.Args[1]
	secretKey = os.Args[2]
	baseURL   = "https://api.binance.com"
)

func generateSignature(params map[string]string) string {
	// Create a slice of parameter keys
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}

	// Sort the keys alphabetically
	sort.Strings(keys)

	// Construct the query string
	var queryString string
	for _, k := range keys {
		queryString += k + "=" + params[k] + "&"
	}

	// Remove the trailing "&"
	queryString = strings.TrimSuffix(queryString, "&")

	// Hash the query string using HMAC SHA256 with the API secret
	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(queryString))
	signature := hex.EncodeToString(mac.Sum(nil))

	return signature
}

func main() {
	client := resty.New()

	endpoint := "/api/v3/account"
	params := map[string]string{
		"timestamp": fmt.Sprintf("%d", time.Now().Unix()*1000),
	}
	signature := generateSignature(params)
	params["signature"] = signature

	responce, err := client.R().
		SetQueryParams(params).
		SetHeader("X-MBX-APIKEY", apiKey).
		Get(baseURL + endpoint)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Responce:", string(responce.Body()))
}
