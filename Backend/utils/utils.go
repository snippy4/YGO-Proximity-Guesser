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

type CardSet struct {
	SetName       string `json:"set_name"`
	SetCode       string `json:"set_code"`
	SetRarity     string `json:"set_rarity"`
	SetRarityCode string `json:"set_rarity_code"`
	SetPrice      string `json:"set_price"`
}

type Card struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	YgoProDeckURL string `json:"ygoprodeck_url"`
}

type CardData struct {
	Data []Card `json:"data"`
}

func GetCard(cardName string) string {
	file, err := os.Open("cards.json")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var cardData CardData
	if err := json.Unmarshal(byteValue, &cardData); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	targetName := cardName
	var found *Card
	for _, card := range cardData.Data {
		if card.Name == targetName {
			found = &card
			break
		}
	}

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
	file, err := os.Open("cards.json")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var cardData CardData
	if err := json.Unmarshal(byteValue, &cardData); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	var matchedCards []Card
	for _, card := range cardData.Data {
		if strings.Contains(strings.ToLower(card.Name), strings.ToLower(cardName)) {
			matchedCards = append(matchedCards, card)
		}
	}

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
	if id == answer {
		return "Correct"
	}
	file, err := os.Open("data/" + answer + " sorted list.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	place := 20
	for scanner.Scan() {
		key, value := strings.Split(scanner.Text(), ":")[0], strings.Split(scanner.Text(), ":")[1]
		ids := strings.Trim(key, "()")
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

	return "No matching value found"
}

func CardByID(id string) string {
	file, err := os.Open("cards.json")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var cardData CardData
	if err := json.Unmarshal(byteValue, &cardData); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	id = strings.TrimSpace(id)
	var matchedCards []Card
	for _, card := range cardData.Data {
		if strconv.Itoa(card.ID) == id {
			matchedCards = append(matchedCards, card)
		}
	}

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

func GetHint(closest string, answer string, failed bool) string {
	if closest == "" {
		return "23434538"
	}
	if closest == answer {
		return answer
	}
	file, err := os.Open("data/" + answer + " sorted list.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	closestPos := 9999
	newPos := 0
	for scanner.Scan() {
		key, _ := strings.Split(scanner.Text(), ":")[0], strings.Split(scanner.Text(), ":")[1]
		ids := strings.Trim(key, "()")
		idParts := strings.Split(ids, ",")
		if len(idParts) == 2 {
			id1 := strings.TrimSpace(idParts[0])
			id1 = strings.ReplaceAll(id1, "'", "")
			id2 := strings.TrimSpace(idParts[1])
			id2 = strings.ReplaceAll(id2, "'", "")
			closest = strings.TrimSpace(closest)
			if closest == id1 || closest == id2 {
				closestPos = newPos
			}
		}
		newPos++
	}
	if closestPos == 9999 {
		return "23434538"
	}
	if closestPos == 0 {
		closestPos = 42
	}
	if !failed {
		newPos = closestPos / 2
	} else {
		newPos = closestPos/2 + 1
	}
	count := 0
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatalf("Error seeking to start of file: %v", err)
	}
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		if count == newPos {
			key, _ := strings.Split(scanner.Text(), ":")[0], strings.Split(scanner.Text(), ":")[1]
			ids := strings.Trim(key, "()")
			idParts := strings.Split(ids, ",")
			if len(idParts) == 2 {
				id1 := strings.TrimSpace(idParts[0])
				id1 = strings.ReplaceAll(id1, "'", "")
				id2 := strings.TrimSpace(idParts[1])
				id2 = strings.ReplaceAll(id2, "'", "")
				if answer == id1 {
					return id2
				} else {
					return id1
				}
			}
		}
		count++
	}
	return ""
}

func ListClosestsCards(answer string) string {
	IDValuePairs := make(map[string]float64, 0)
	file, err := os.Open("data/" + answer + " sorted list.txt")
	if err != nil {
		log.Fatalf("Error seeking to start of file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() && count < 50 {
		key, value := strings.Split(scanner.Text(), ":")[0], strings.Split(scanner.Text(), ":")[1]
		ids := strings.Trim(key, "()")
		idParts := strings.Split(ids, ",")
		if len(idParts) == 2 {
			value = strings.ReplaceAll(value, " ", "")
			if count < 20 {
				value = strconv.Itoa(20 - count)
			}
			id1 := strings.TrimSpace(idParts[0])
			id1 = strings.ReplaceAll(id1, "'", "")
			id2 := strings.TrimSpace(idParts[1])
			id2 = strings.ReplaceAll(id2, "'", "")
			if answer == id1 {
				IDValuePairs[id2], err = strconv.ParseFloat(value, 64)
			} else {
				IDValuePairs[id1], err = strconv.ParseFloat(value, 64)
			}
		}
		count++
	}
	response, err := json.Marshal(IDValuePairs)
	return string(response)
}
