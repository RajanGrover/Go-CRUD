package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Item struct {
	ItemId string `json:id`
	Name   string `json:name`
	Price  int    `json:price`
}

var items []Item

func getAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	// the response that we r sending we need to encode in json
	json.NewEncoder(w).Encode(items)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r) /// here we will be getting the id from the request
	for index, item := range items {
		if item.ItemId == params["id"] { // find the item and delete it
			items = append(items[:index], items[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(items)
}
func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range items {
		if item.ItemId == params["id"] { // find the item encode it and send it that is return it
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// json.NewEncoder(w).Encode(&Item{}) // doubt

}
func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ItemId = strconv.Itoa(rand.Intn(10000))
	items = append(items, item)
	json.NewEncoder(w).Encode(items)
}

// easy method just lets make the combo of delete +create
func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r) // get the id
	for index, item := range items {
		if item.ItemId == params["id"] { // find the item encode it and send it that is return it
			items = append(items[:index], items[index+1:]...) // delete the one that is to be updated
			var item Item                                     // now adding a new one
			_ = json.NewDecoder(r.Body).Decode(&item)
			item.ItemId = strconv.Itoa(rand.Intn(10000))
			items = append(items, item)
			json.NewEncoder(w).Encode(items)
			return
		}
	}
}

func main() {

	items = append(items, Item{"1", "Snacks", 100})
	items = append(items, Item{"2", "Chips", 200})

	r := mux.NewRouter()
	r.HandleFunc("/items", getAllItems).Methods("GET")
	r.HandleFunc("/items/{id}", getItem).Methods("GET")
	r.HandleFunc("/items", createItem).Methods("POST")
	r.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	r.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")
	fmt.Println("Starting server at port :8000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
