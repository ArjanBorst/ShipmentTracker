package main

//https://www.jsonschemavalidator.net/

import (
	"encoding/json"
	"fmt"
	sApi "main/ship24/api"
	"net/http"
)

func HTTPSDelivered(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Println("Endpoint Hit: GetJsonStatistics")

	mux.RLock()
	json.NewEncoder(w).Encode(mDelivered)
	mux.RUnlock()
}

func HTTPPending(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Println("Endpoint Hit: GetJsonStatistics")

	mux.RLock()
	json.NewEncoder(w).Encode(mPending)
	mux.RUnlock()
}

func HTTPSNotShipped(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Println("Endpoint Hit: GetJsonStatistics")

	mux.RLock()
	json.NewEncoder(w).Encode(mNotShipped)
	mux.RUnlock()
}

func HTTPGetCouriers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Println("Endpoint Hit: GetJsonStatistics")

	mux.RLock()
	json.NewEncoder(w).Encode(sApi.Courier)
	mux.RUnlock()
}

func HTTPGetCountries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Println("Endpoint Hit: GetJsonStatistics")

	mux.RLock()
	json.NewEncoder(w).Encode(sApi.Courier)
	mux.RUnlock()
}
