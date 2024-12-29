package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type kv struct {
	Key   string
	Value float64
}

func Sorted_list(input string) {
	file, err := os.Open("proximity.json")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	var data map[string]float64
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	filtered := make(map[string]float64)
	for key, value := range data {
		if strings.Contains(key, input) {
			filtered[key] = value
		}
	}
	w_max := 0.0
	w_min := math.MaxFloat64
	for _, value := range filtered {
		if value > w_max {
			w_max = value
		}
		if value < w_min {
			w_min = value
		}
	}
	w_max = math.Log(w_max)
	w_min = math.Log(w_min)

	var sorted []kv
	for k, v := range filtered {
		sorted = append(sorted, kv{Key: k, Value: (math.Log(v) - w_min) / (w_max - w_min)})
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value < sorted[j].Value
	})

	newfile, err := os.Create("data/" + input + " sorted list.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer newfile.Close()

	writer := bufio.NewWriter(newfile)
	lines := len(sorted) - 1
	for i, _ := range sorted {
		item := sorted[lines-i]
		fmt.Printf("%s: %.2f\n", item.Key, item.Value)
		outline := fmt.Sprintf("%s: %.2f\n", item.Key, item.Value)
		_, err := writer.WriteString(outline)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Filtered map written to filtered_output.json")

}

func CleanSortedList(id string) {
	file, err := os.Open("data/" + id + " sorted list.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	cardsjson, err := os.ReadFile("cards.json")
	if err != nil {
		log.Fatalf("Error openening file %v", err)
	}
	cards := string(cardsjson)
	defer file.Close()
	var sorted []kv
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		key, value := strings.Split(scanner.Text(), ":")[0], strings.Split(scanner.Text(), ":")[1]
		ids := strings.Trim(key, "()")
		idParts := strings.Split(ids, ",")
		if len(idParts) == 2 {
			id1 := strings.TrimSpace(idParts[0])
			id1 = strings.ReplaceAll(id1, "'", "")
			id2 := strings.TrimSpace(idParts[1])
			id2 = strings.ReplaceAll(id2, "'", "")
			floatvalue, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
			if err != nil {
				log.Fatalf("Error parsing float: %v", err)
			}
			if strings.Contains(cards, id1+",\n      \"name\":") && strings.Contains(cards, id2+",\n      \"name\":") {
				sorted = append(sorted, kv{Key: key, Value: floatvalue})
			}
		}
	}
	newfile, err := os.Create("data/" + id + " sorted list.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer newfile.Close()

	writer := bufio.NewWriter(newfile)
	for i, _ := range sorted {
		item := sorted[i]
		fmt.Printf("%s: %.2f\n", item.Key, item.Value)
		outline := fmt.Sprintf("%s: %.2f\n", item.Key, item.Value)
		_, err := writer.WriteString(outline)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Filtered map written to filtered_output.json")

}
