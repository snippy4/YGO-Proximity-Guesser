package main

import (
	"Backend/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Data represents a structure for JSON responses
type Data struct {
	Message string `json:"message"`
}

func main() {
	// Serve static files (HTML, CSS, JS)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// API endpoint
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := Data{Message: utils.GetCard("Performage Plushfire")}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Start the server
	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
