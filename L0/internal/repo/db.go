package repo

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"wb_l0/internal/entity"
)

var db *sql.DB

func InitDB() {
	var err error
	connStr := "user=postgres password= dbname=orders_db2 sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func SaveOrderToDB(order entity.Order) {
	deliveryJSON, err := json.Marshal(order.Delivery)
	if err != nil {
		log.Println("Error marshaling delivery:", err)
		return
	}

	paymentJSON, err := json.Marshal(order.Payment)
	if err != nil {
		log.Println("Error marshaling payment:", err)
		return
	}

	itemsJSON, err := json.Marshal(order.Items)
	if err != nil {
		log.Println("Error marshaling items:", err)
		return
	}

	query := `INSERT INTO orders(
        order_uid, track_number, entry, delivery, payment, items, locale, internal_signature, customer_id, 
        delivery_service, shardkey, sm_id, date_created, oof_shard) 
        VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
        ON CONFLICT (order_uid) DO NOTHING;`

	_, err = db.Exec(query, order.OrderUID, order.TrackNumber, order.Entry, deliveryJSON, paymentJSON, itemsJSON, order.Locale,
		order.InternalSignature, order.CustomerID, order.DeliveryService, order.Shardkey, order.SmID, order.DateCreated, order.OofShard)

	if err != nil {
		log.Println("Error inserting order into DB:", err)
	}
}
