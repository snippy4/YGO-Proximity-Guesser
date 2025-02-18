package main

import (
	"Backend/utils"
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

type SuggestionResponse struct {
	Suggestions []string `json:"suggestions"`
}

type SelectResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	newRandomCard()
	//current_daily = "23434538"
	//fmt.Println(utils.CardTypeByID(current_daily))
	// testing code commented out
	go startHTTPServer()
	go dailyReset()
	select {}
}

func startHTTPServer() {
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		suggestions := getSuggestions(query)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(suggestions)
	})

	http.HandleFunc("/select", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		result := getResult(query)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	http.HandleFunc("/hint", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		result := getHint(query)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	http.HandleFunc("/answer", func(w http.ResponseWriter, r *http.Request) {
		result := utils.CardByID(current_daily)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		result := getList()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	http.HandleFunc("/bonus", func(w http.ResponseWriter, r *http.Request) {
		result := getBonus()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	http.ListenAndServe(":8080", nil)
}

func getSuggestions(query string) []string {
	if len(query) > 1000 {
		return make([]string, 0)
	}
	cardsList := make([]string, 0)
	cards := utils.SearchCards(query)
	cardsList = append(cardsList, cards)
	return cardsList
}

func getResult(query string) map[string]string {
	mu.Lock()
	cardsJSON := utils.SearchCards(query)

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
	utils.CleanSortedList(new_card)
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

func getHint(q string) string {
	mu.Lock()
	hint := utils.GetHint(q, current_daily, false)
	hintJSON := utils.CardByID(hint)
	fmt.Println(hintJSON)
	var cards []utils.Card
	err := json.Unmarshal([]byte(hintJSON), &cards)
	if err != nil {
		fmt.Println(err)
	}
	hintName := ""
	if len(cards) == 0 {
		fmt.Println("failed to find card: " + hint)
		hint = utils.GetHint(q, current_daily, true)
		hintJSON = utils.CardByID(hint)
		fmt.Println(hintJSON)
		err = json.Unmarshal([]byte(hintJSON), &cards)
		if err != nil {
			fmt.Println(err)
		}
		if len(cards) == 0 {
			return ""
		} else {
			hintName = cards[0].Name
		}
	} else {
		hintName = cards[0].Name
	}
	mu.Unlock()
	return hintName
}

func getList() string {
	mu.Lock()
	list := utils.ListClosestsCards(current_daily)
	mu.Unlock()
	return list
}

func getBonus() string {
	mu.Lock()
	value := utils.CardTypeByID(current_daily)
	mu.Unlock()
	return value
}
