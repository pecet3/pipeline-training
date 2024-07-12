package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Metoda nie dozwolona", http.StatusMethodNotAllowed)
		return
	}

	now := time.Now()
	response := Response{
		Date: now.Format("2006-01-02"),
		Time: now.Format("15:04:05"),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", handleRequest)

	err := http.ListenAndServe(":8070", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
