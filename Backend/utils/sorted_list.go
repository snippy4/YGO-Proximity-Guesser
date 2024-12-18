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
	"strings"
)

func Sorted_list(input string) {
	// Read the JSON file
	file, err := os.Open("proximity.json")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Parse the JSON into a map
	var data map[string]float64
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	// Filter the map to include only keys containing "id"
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
	// Convert the filtered map to a slice of key-value pairs for sorting
	type kv struct {
		Key   string
		Value float64
	}

	var sorted []kv
	for k, v := range filtered {
		sorted = append(sorted, kv{Key: k, Value: (math.Log(v) - w_min) / (w_max - w_min)})
	}

	// Sort the slice by values
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value < sorted[j].Value
	})

	newfile, err := os.Create("data/" + input + " sorted list.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer newfile.Close()

	// Print the sorted results
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
