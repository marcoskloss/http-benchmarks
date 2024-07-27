package main

import (
	"net/http"
	"encoding/json"
	"log"
)

type Response struct {
	Ok	bool `json:"ok"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Ok: true})
	})

  log.Println("go/standard listening on port 3003");
	log.Fatal(http.ListenAndServe(":3003", nil))
}