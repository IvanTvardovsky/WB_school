package main

import (
	"encoding/json"
	"log"
	"time"
	"wb_l0/internal/entity"

	stan "github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("test-cluster", "test-publisher")
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	order := entity.Order{
		OrderUID:    "b563feb7b2b84b6test1111",
		TrackNumber: "WBILMTESTTRACK",
		Entry:       "WBIL",
		Delivery: entity.Delivery{
			Name:    "Test Testov",
			Phone:   "+9720000000",
			Zip:     "2639809",
			City:    "Kiryat Mozkin",
			Address: "Ploshad Mira 15",
			Region:  "Kraiot",
			Email:   "test@gmail.com",
		},
		Payment: entity.Payment{
			Transaction:  "b563feb7b2b84b6test",
			RequestID:    "",
			Currency:     "USD",
			Provider:     "wbpay",
			Amount:       1817,
			PaymentDt:    1637907727,
			Bank:         "alpha",
			DeliveryCost: 1500,
			GoodsTotal:   317,
			CustomFee:    0,
		},
		Items: []entity.Item{
			{
				ChrtID:      9934930,
				TrackNumber: "WBILMTESTTRACK",
				Price:       453,
				Rid:         "ab4219087a764ae0btest",
				Name:        "Mascaras",
				Sale:        30,
				Size:        "0",
				TotalPrice:  317,
				NmID:        2389212,
				Brand:       "Vivienne Sabo",
				Status:      202,
			},
		},
		Locale:            "en",
		InternalSignature: "",
		CustomerID:        "test",
		DeliveryService:   "meest",
		Shardkey:          "9",
		SmID:              99,
		DateCreated:       time.Now(),
		OofShard:          "1",
	}

	orderData, err := json.Marshal(order)
	if err != nil {
		log.Fatal("Error marshaling order:", err)
	}

	err = sc.Publish("orders", orderData)
	if err != nil {
		log.Fatal("Error publishing order:", err)
	}

	log.Println("Test order published to NATS")
}
