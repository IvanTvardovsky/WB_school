package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"wb_l0/internal/repo"
)

func StartHTTPServer() {
	r := mux.NewRouter()
	r.HandleFunc("/order/{id}", GetOrderHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["id"]

	order, found := repo.GetOrderFromCache(orderID)
	if !found {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	err := json.NewEncoder(w).Encode(order)
	if err != nil {
		http.Error(w, "Can not encode", http.StatusInternalServerError)
		return
	}
}
