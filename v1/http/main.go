package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Status string          `json:"status"`
	Data   json.RawMessage `json:"data"`
}

func main() {
	http.HandleFunc("/", GetHandler)

	addr := "localhost:8050"

	log.Println("===================================")
	log.Println("Fehmi Local HTTP Server Started")
	log.Printf("Listening on http://%s\n", addr)
	log.Println("===================================")

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	data, err := os.ReadFile("D:\\Projects\\backend\\agent\\v1\\http\\data.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !json.Valid(data) {
		http.Error(w, "Invalid JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(Response{
		Status: "success",
		Data:   json.RawMessage(data),
	}); err != nil {
		log.Printf("Response encode error: %v", err)
	}
}
