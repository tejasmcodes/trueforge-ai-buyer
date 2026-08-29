package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/tejasmcodes/trueforge-ai-buyer/internal/mcpserver"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
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

	mcpServer := mcpserver.NewServer()
	mcpHandler := mcp.NewStreamableHTTPHandler(
		func(*http.Request) *mcp.Server {
			return mcpServer
		},
		nil,
	)

	mux.Handle("/mcp", mcpHandler)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}
}
