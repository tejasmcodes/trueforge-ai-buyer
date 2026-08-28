package main

import (
	"encoding/json"
	"net/http"
	"log"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("TrueForge Agent Harness Project"))
	})
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		message := Message{
			Message: "Health is OK",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(message)
	})

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Failed to start the server!")
	}

}
