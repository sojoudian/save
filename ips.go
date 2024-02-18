package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Define a struct to match the JSON structure
type logEntry struct {
	IP   string `json:"ip"`
	Time string `json:"time"`
}

func main() {
	// Open the file
	file, err := os.Open("ips.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Use a map to track unique IPs
	uniqueIPs := make(map[string]bool)

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Unmarshal the JSON line into the logEntry struct
		var entry logEntry
		if err := json.Unmarshal(scanner.Bytes(), &entry); err != nil {
			fmt.Println("Error parsing line:", err)
			continue
		}

		// Add the IP to the map (automatically handles uniqueness)
		uniqueIPs[entry.IP] = true
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	// Print the unique IPs
	fmt.Println("Unique IPs:")
	for ip := range uniqueIPs {
		fmt.Println(ip)
	}
}
