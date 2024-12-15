package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("proximity.json")
	if err != nil {
		fmt.Println("err")
		return
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
	current_node := ""
	i := 1
	for node, _ := range data {
		if i == 1 {
			current_node = node
		}
	}
	current_node = strings.Split(strings.Trim(current_node, "()"), ",")[0]
	current_node = strings.ReplaceAll(current_node, "'", "")
	degree := make(map[string]int)

	// Iterate over the edges
	for edge, _ := range data {
		u, v := strings.ReplaceAll(strings.Split(strings.Trim(edge, "()"), ",")[0], "'", ""), strings.ReplaceAll(strings.Split(strings.Trim(edge, "()"), ",")[1], "'", "")
		degree[u]++
		degree[v]++
	}
	fmt.Println(degree)
	fmt.Println(degree[current_node])

}
