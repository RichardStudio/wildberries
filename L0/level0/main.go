package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"level0/database"
	"level0/handlers"
	"level0/kafka"
)

func CreateDbConfig() database.DatabaseConfig {
	return database.DatabaseConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "20042001",
		DBName:   "orders",
	}
}

func main() {
	client, err := database.NewClient(CreateDbConfig())
	if err != nil {
		log.Fatal(err)
	}
	defer client.DBClose()

	client.RestoreCacheFromDB()

	go kafka.ConsumeKafkaMessages(client)

	r := mux.NewRouter()
	r.HandleFunc("/orders/{id}", handlers.GetOrder).Methods("GET")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
