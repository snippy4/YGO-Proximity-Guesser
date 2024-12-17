package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"strings"
)

func weightedQuadratic(x, max float64) float64 {
	center := 0.55 * max
	scale := 0.75 * max * 0.6

	// Quadratic formula
	return math.Max(0, 1-math.Pow((x-center)/scale, 2))
}

func Random_node() (node string) {
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
	current_node = strings.Split(strings.Trim(current_node, "() "), ",")[0]
	current_node = strings.ReplaceAll(current_node, "'", "")
	degree := make(map[string]float64)

	// Iterate over the edges
	for edge, _ := range data {
		u, v := strings.ReplaceAll(strings.Split(strings.Trim(edge, "() "), ",")[0], "'", ""), strings.ReplaceAll(strings.Split(strings.Trim(edge, "()"), ",")[1], "'", "")
		degree[u]++
		degree[v]++
	}
	max_deg := 0.0
	for _, deg := range degree {
		if deg > max_deg {
			max_deg = deg
		}
	}
	for node, deg := range degree {
		prob := weightedQuadratic(deg, max_deg)
		if prob != 0 && prob < 0.7 {
			if rand.Float64() < prob {
				current_node = node
				break
			}
		}
	}

	// Metropolis walk
	fmt.Println(current_node)
	for i := 0; i < 30; i++ {
		edges := make([]string, 0)
		for edge, _ := range data {
			u, v := strings.ReplaceAll(strings.Split(strings.Trim(edge, "() "), ",")[0], "'", ""), strings.ReplaceAll(strings.Split(strings.Trim(edge, "()"), ",")[1], "'", "")
			if u == current_node || v == current_node {
				edges = append(edges, edge)
			}
		}
		for i, _ := range edges {
			candidate_edge := rand.Intn(len(edges) - 1 - i)
			u, v := strings.ReplaceAll(strings.Split(strings.Trim(edges[candidate_edge], "() "), ",")[0], "'", ""), strings.ReplaceAll(strings.Split(strings.Trim(edges[candidate_edge], "()"), ",")[1], "'", "")
			next := ""
			if u == current_node {
				next = v
			} else {
				next = u
			}
			prob := weightedQuadratic(degree[next], max_deg)
			randprob := rand.Float64()
			if randprob < prob {
				current_node = next
				break
			}
		}
	}
	current_node = strings.ReplaceAll(current_node, " ", "")
	fmt.Println(current_node)
	fmt.Println(weightedQuadratic(degree[current_node], max_deg))
	return current_node
}
