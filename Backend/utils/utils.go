package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
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
		return "No such card"
	}
	response, _ := json.Marshal(*found)
	return string(response)
}

func SearchCards(cardName string) string {
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

	// Collect all cards containing cardName in their name
	var matchedCards []Card
	for _, card := range cardData.Data {
		if strings.Contains(strings.ToLower(card.Name), strings.ToLower(cardName)) {
			matchedCards = append(matchedCards, card)
		}
	}

	// Return the result
	if len(matchedCards) > 0 {
		response, err := json.Marshal(matchedCards)
		if err != nil {
			log.Fatalf("Error marshalling matched cards: %v", err)
		}
		return string(response)
	} else {
		return "[]"
	}
}

func FindValueByID(id string, answer string) string {
	// Open the JSON file
	if id == answer {
		return "CARD FOUND WOOO 🎉"
	}
	file, err := os.Open(answer + " sorted list.txt") // Replace with your JSON file name
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Search for the ID in the keys
	place := 20
	for scanner.Scan() {
		key, value := strings.Split(scanner.Text(), ":")[0], strings.Split(scanner.Text(), ":")[1]
		ids := strings.Trim(key, "()") // Remove parentheses from the key
		idParts := strings.Split(ids, ",")
		if len(idParts) == 2 {
			id1 := strings.TrimSpace(idParts[0])
			id1 = strings.ReplaceAll(id1, "'", "")
			id2 := strings.TrimSpace(idParts[1])
			id2 = strings.ReplaceAll(id2, "'", "")
			if id == id1 || id == id2 {
				if place >= 1 {
					return strconv.Itoa(place)
				}
				return value
			}
		}
		place--
	}

	// If no match is found
	return "No matching value found"
}
