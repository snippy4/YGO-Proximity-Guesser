package main

import (
	"Backend/utils"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	current_daily string
	mu            sync.Mutex
)

// Response structures
type SuggestionResponse struct {
	Suggestions []string `json:"suggestions"`
}

type SelectResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	newRandomCard()
	go startHTTPSServer()
	go dailyReset()
	select {}
}

func startHTTPSServer() {
	http.HandleFunc("/search", corsHandler(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		suggestions := getSuggestions(query) // Replace with your logic
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(suggestions)
	}))

	http.HandleFunc("/select", corsHandler(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		result := getResult(query) // Replace with your logic
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}))

	// Configure HTTPS with TLS
	server := &http.Server{
		Addr: ":8080",
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}

	fmt.Println("Starting HTTPS server on port 443")
	err := server.ListenAndServeTLS("/etc/letsencrypt/live/ygoserver.ddns.net/fullchain.pem", "/etc/letsencrypt/live/ygoserver.ddns.net/privkey.pem")
	if err != nil {
		fmt.Printf("HTTPS server failed: %s\n", err)
	}
}

// CORS Middleware
func corsHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins, change "*" to specific domain for stricter security
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler(w, r)
	}
}

// Replace this with your logic to generate suggestions
func getSuggestions(query string) []string {
	if len(query) > 1000 {
		return make([]string, 0)
	}
	cardsList := make([]string, 0)
	cards := utils.SearchCards(query)
	cardsList = append(cardsList, cards)
	return cardsList
}

// Replace this with your logic to return a key-value pair
func getResult(query string) map[string]string {
	mu.Lock()
	cardsJSON := utils.SearchCards(query) // Call the modified function with query input

	var cards []utils.Card
	err := json.Unmarshal([]byte(cardsJSON), &cards)
	if err != nil {
		fmt.Println(err)
	}
	var card utils.Card
	for _, cardselect := range cards {
		if cardselect.Name == query {
			card = cardselect
		}
	}
	key := card.Name
	value := utils.FindValueByID(strconv.Itoa(card.ID), current_daily)
	mu.Unlock()
	return map[string]string{"key": key, "value": value, "id": strconv.Itoa(card.ID)}
}

func newRandomCard() {
	mu.Lock()
	new_card := utils.Random_node()
	utils.Sorted_list(new_card)
	current_daily = new_card
	mu.Unlock()
}

func dailyReset() {
	for {
		now := time.Now().UTC()
		nextReset := time.Date(now.Year(), now.Month(), now.Day(), 6, 0, 0, 0, time.UTC)
		if nextReset.Before(now) {
			nextReset = nextReset.Add(24 * time.Hour)
		}

		durationUntilReset := time.Until(nextReset)
		fmt.Printf("Next daily reset scheduled in: %s\n", durationUntilReset)
		time.Sleep(durationUntilReset)
		newRandomCard()
	}
}
