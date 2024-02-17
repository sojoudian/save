package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/save", saveIPHandler)
	log.Println("Server started on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*") // Adjust the allowed origin as needed
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func saveIPHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	switch r.Method {
	case "POST":
		enableCors(&w)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		if r.Method != "POST" {
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusOK)
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		var data map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Error parsing JSON request body", http.StatusBadRequest)
			return
		}

		// Add the current time to the data map
		currentTime := time.Now().Format("2006-01-02 15:04:05") // Use Go's reference time format
		data["time"] = currentTime

		// Convert the JSON data with time added back to a string to save to the file
		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Error converting JSON data to string", http.StatusInternalServerError)
			return
		}

		// Append the JSON string to the file
		file, err := os.OpenFile("ips.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			http.Error(w, "Error opening file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		if _, err := file.WriteString(string(jsonData) + "\n"); err != nil {
			http.Error(w, "Error writing to file", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Data saved successfully")
		fmt.Println(string(jsonData), "Data saved successfully")

	case "GET":
		// Implement GET method functionality if needed
		fmt.Fprintf(w, "Hello World\n")

	default:
		http.Error(w, "Only GET and POST methods are allowed", http.StatusMethodNotAllowed)
	}
}

