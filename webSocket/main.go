package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	endpoint = "wss://stream.binance.com:9443/ws/your_api_key/myTrades"
)

type Trade struct {
	Symbol          string `json:"symbol"`
	Id              int    `json:"id"`
	OrderId         int    `json:"orderId"`
	Price           string `json:"price"`
	Quantity        string `json:"qty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            int64  `json:"time"`
	IsBuyer         bool   `json:"isBuyer"`
	IsMaker         bool   `json:"isMaker"`
	IsBestMatch     bool   `json:"isBestMatch"`
}

func main() {
	// Establish a connection
	conn, _, err := websocket.DefaultDialer.Dial(endpoint, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	// Read messages from the connection
	var trades []Trade
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}

			var trade Trade
			if err := json.Unmarshal(message, &trade); err != nil {
				log.Println("json unmarshal error:", err)
				continue
			}
			trades = append(trades, trade)
		}
	}()

	// Subscribe to the 'myTrades' channel
	// Note: Adjust the subscription message as required by Binance
	// conn.WriteMessage(websocket.TextMessage, []byte(`{"method":"SUBSCRIBE","params":["your_api_key@myTrades"],"id":1}`))

	// Periodically ping the server to keep the connection alive
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	// Run for some time to fetch trades
	time.Sleep(5 * time.Minute)

	// Stop the connection
	if err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
		log.Println("write close:", err)
		return
	}
	select {
	case <-done:
	case <-time.After(time.Second):
	}
	log.Println("Trades:", trades)

	// Convert trades to JSON
	tradesJSON, err := json.Marshal(trades)
	if err != nil {
		log.Println("json marshal error:", err)
		return
	}
	log.Println("Trades in JSON format:", string(tradesJSON))
}
