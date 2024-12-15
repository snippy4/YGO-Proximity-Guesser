package main

import (
	"Backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	go startHTTPServer()

	select {}
}

func startHTTPServer() {
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		suggestions := getSuggestions(query) // Replace with your logic
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(suggestions)
	})

	http.HandleFunc("/select", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		result := getResult(query) // Replace with your logic
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	http.ListenAndServe(":8080", nil)
}

// Replace this with your logic to generate suggestions
func getSuggestions(query string) []string {
	cardsList := make([]string, 0)
	cards := utils.SearchCards(query)
	cardsList = append(cardsList, cards)
	return cardsList
}

// Replace this with your logic to return a key-value pair
func getResult(query string) map[string]string {
	cardsJSON := utils.SearchCards(query) // Call the modified function with query input

	var cards []utils.Card
	err := json.Unmarshal([]byte(cardsJSON), &cards)
	if err != nil {
		fmt.Println(err)
	}
	key := cards[0].Name
	value := utils.FindValueByID(strconv.Itoa(cards[0].ID))
	return map[string]string{"key": key, "value": value}
}
