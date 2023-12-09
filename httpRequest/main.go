package main

import (
	"fmt"
	"os"

	"github.com/go-resty/resty"
)

var (
	apiKey    = os.Args[1]
	secretKey = os.Args[2]
	baseURL   = "https://api.binance.com"
)

func main() {
	client := resty.New()

	responce, err := client.R().
		SetHeader("X-MBX-APIKEY", apiKey).
		Get(baseURL + "/api/v3/account")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Responce:", string(responce.Body()))
}
