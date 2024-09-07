package repo

import (
	"encoding/json"
	"log"
	"sync"
	"wb_l0/internal/entity"
)

var orderCache = make(map[string]entity.Order)
var cacheLock sync.RWMutex

func SaveOrderToCache(order entity.Order) {
	cacheLock.Lock()
	defer cacheLock.Unlock()
	orderCache[order.OrderUID] = order
}

func GetOrderFromCache(orderID string) (entity.Order, bool) {
	cacheLock.RLock()
	defer cacheLock.RUnlock()
	order, found := orderCache[orderID]
	return order, found
}

func RestoreCacheFromDB() {
	rows, err := db.Query("SELECT order_uid, track_number, entry, delivery, payment, items, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard FROM orders")
	if err != nil {
		log.Fatal("Error restoring cache from DB:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var order entity.Order
		var deliveryJSON, paymentJSON, itemsJSON []byte

		err := rows.Scan(
			&order.OrderUID,
			&order.TrackNumber,
			&order.Entry,
			&deliveryJSON,
			&paymentJSON,
			&itemsJSON,
			&order.Locale,
			&order.InternalSignature,
			&order.CustomerID,
			&order.DeliveryService,
			&order.Shardkey,
			&order.SmID,
			&order.DateCreated,
			&order.OofShard,
		)
		if err != nil {
			log.Println("Error scanning order:", err)
			continue
		}

		if err = json.Unmarshal(deliveryJSON, &order.Delivery); err != nil {
			log.Println("Error unmarshaling delivery:", err)
			continue
		}

		if err = json.Unmarshal(paymentJSON, &order.Payment); err != nil {
			log.Println("Error unmarshaling payment:", err)
			continue
		}

		if err = json.Unmarshal(itemsJSON, &order.Items); err != nil {
			log.Println("Error unmarshaling items:", err)
			continue
		}

		SaveOrderToCache(order)
	}

	if err := rows.Err(); err != nil {
		log.Fatal("Error iterating through orders:", err)
	}

	log.Println("Cache restored from DB")
}
