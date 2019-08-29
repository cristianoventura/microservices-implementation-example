package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", SayHello)
	log.Fatal(http.ListenAndServe(":3000", router))
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"hello":        "world",
		"from":         "a service written in golang",
		"message sent": "",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
