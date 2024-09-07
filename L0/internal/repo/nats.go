package repo

import (
	"encoding/json"
	stan "github.com/nats-io/stan.go"
	"log"
	"wb_l0/internal/entity"
)

func SubscribeToNATS() {
	sc, err := stan.Connect("test-cluster", "order-service")
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	_, err = sc.Subscribe("orders", func(m *stan.Msg) {
		var order entity.Order
		err := json.Unmarshal(m.Data, &order)
		if err != nil {
			log.Println("Invalid order format:", err)
			return
		}

		SaveOrderToCache(order)
		SaveOrderToDB(order)
	}, stan.DeliverAllAvailable())

	if err != nil {
		log.Fatal(err)
	}

	select {}
}
