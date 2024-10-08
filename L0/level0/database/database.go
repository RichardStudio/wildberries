package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"level0/models"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

type DatabaseClient struct {
	db *sql.DB
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

var Cache = make(map[string]models.Order)
var once sync.Once
var dbError error
var dbClient *DatabaseClient

func NewClient(config DatabaseConfig) (*DatabaseClient, error) {
	once.Do(func() {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.Host, config.Port, config.User, config.Password, config.DBName)
		DB, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			dbError = err
			return
		}
		err = DB.Ping()
		if err != nil {
			dbError = err
			return
		}
		dbClient = &DatabaseClient{
			db: DB,
		}
	})
	if dbError != nil {
		return nil, dbError
	}
	log.Println("Successfully connected to database")
	return dbClient, nil
}

func (c *DatabaseClient) DBClose() {
	err := c.db.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully closed database")
}

func (c *DatabaseClient) RestoreCacheFromDB() {
	if c.db == nil {
		log.Fatal("Database connection is not initialized")
	}
	rows, err := c.db.Query("SELECT * FROM orders")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var order models.Order
		var delivery, payment, items []byte
		if err := rows.Scan(&order.OrderUID, &order.TrackNumber, &order.Entry, &delivery, &payment,
			&items, &order.Locale, &order.InternalSignature, &order.CustomerID, &order.DeliveryService,
			&order.Shardkey, &order.SmID, &order.DateCreated, &order.OofShard); err != nil {
			log.Fatal(err)
			return
		}
		err = json.Unmarshal(delivery, &order.Delivery)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = json.Unmarshal(payment, &order.Payment)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = json.Unmarshal(items, &order.Items)
		if err != nil {
			log.Fatal(err)
			return
		}

		Cache[order.OrderUID] = order
		log.Println("successfully restored from db")
	}
}

func (c *DatabaseClient) SaveOrder(order models.Order) {

	delivery, err := json.Marshal(order.Delivery)
	if err != nil {
		log.Printf("Error marshalling delivery: %v", err)
		return
	}

	payment, err := json.Marshal(order.Payment)
	if err != nil {
		log.Printf("Error marshalling payment: %v", err)
		return
	}

	items, err := json.Marshal(order.Items)
	if err != nil {
		log.Printf("Error marshalling items: %v", err)
		return
	}
	_, err = c.db.Exec(`INSERT INTO orders (order_uid, track_number, entry, delivery, payment, items,
                    locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) 
					ON CONFLICT (order_uid) DO 
					UPDATE SET track_number = EXCLUDED.track_number, entry = EXCLUDED.entry, 
					    delivery = EXCLUDED.delivery, payment = EXCLUDED.payment, items = EXCLUDED.items, 
					    locale = EXCLUDED.locale, internal_signature = EXCLUDED.internal_signature, 
					    customer_id = EXCLUDED.customer_id, delivery_service = EXCLUDED.delivery_service, 
					    shardkey = EXCLUDED.shardkey, sm_id = EXCLUDED.sm_id, date_created = EXCLUDED.date_created, 
					    oof_shard = EXCLUDED.oof_shard`,
		order.OrderUID, order.TrackNumber, order.Entry, delivery, payment, items, order.Locale,
		order.InternalSignature, order.CustomerID, order.DeliveryService, order.Shardkey, order.SmID,
		order.DateCreated, order.OofShard)
	if err != nil {
		log.Printf("Error inserting/updating order in DB: %v", err)
	}
	log.Println("successfully inserted order into DB")
}
