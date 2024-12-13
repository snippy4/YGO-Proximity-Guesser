package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// CardSet represents the card set details
type CardSet struct {
	SetName       string `json:"set_name"`
	SetCode       string `json:"set_code"`
	SetRarity     string `json:"set_rarity"`
	SetRarityCode string `json:"set_rarity_code"`
	SetPrice      string `json:"set_price"`
}

// Card represents a single card's details
type Card struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	YgoProDeckURL string `json:"ygoprodeck_url"`
}

// CardData represents the top-level structure
type CardData struct {
	Data []Card `json:"data"`
}

func GetCard(cardName string) string {
	// Open the JSON file
	file, err := os.Open("cards.json")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Read the file contents
	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Parse the JSON into the CardData struct
	var cardData CardData
	if err := json.Unmarshal(byteValue, &cardData); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Search for a card by name
	targetName := cardName
	var found *Card
	for _, card := range cardData.Data {
		if card.Name == targetName {
			found = &card
			break
		}
	}

	// Print the result
	if found != nil {
		fmt.Printf("Found card: %+v\n", *found)
	} else {
		fmt.Printf("Card with name '%s' not found.\n", targetName)
	}
	response, err := json.Marshal(*found)
	return string(response)
}
