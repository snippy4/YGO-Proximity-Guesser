package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	input := "26077387"
	// Read the JSON file
	file, err := os.Open("proximity.json")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Parse the JSON into a map
	var data map[string]float64
	bytes, err := ioutil.ReadAll(file)
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

	// Convert the filtered map to a slice of key-value pairs for sorting
	type kv struct {
		Key   string
		Value float64
	}

	var sorted []kv
	for k, v := range filtered {
		sorted = append(sorted, kv{Key: k, Value: v})
	}

	// Sort the slice by values
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value < sorted[j].Value
	})

	// Print the sorted results
	fmt.Println("Sorted results:")
	for _, item := range sorted {
		fmt.Printf("%s: %.2f\n", item.Key, item.Value)
	}

	newfile, err := os.Create(input + " sorted list.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer newfile.Close()

	jsonData, err := json.MarshalIndent(filtered, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	_, err = newfile.Write(jsonData)
	if err != nil {
		log.Fatalf("Failed to write JSON to file: %v", err)
	}

	fmt.Println("Filtered map written to filtered_output.json")

}
