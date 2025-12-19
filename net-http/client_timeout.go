package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	
	// We are connecting to a site that deliberately delays the response.
	// "httpbin.org/delay/5" waits 5 seconds before replying.
	slowURL := "https://httpbin.org/delay/5"

	fmt.Println("--- Setting up the Client with a Stopwatch ---")
	
	// CONSTRUCT: http.Client
	// Instead of using the default http.Get, we build a custom "Client".
	// This allows us to set rules, like a Timeout.
	client := http.Client{
		Timeout: 2 * time.Second,
	}

	fmt.Printf("Attempting to fetch data from %s...\n", slowURL)
	fmt.Println("(I will give up in 2 seconds)")
	
	// Start the timer
	start := time.Now()

	// METHOD: client.Get(url)
	// This works exactly like http.Get, but it follows the rules we set above.
	resp, err := client.Get(slowURL)

	// Calculate how long we waited
	duration := time.Since(start)

	if err != nil {
		// We EXPECT an error here because 5s (server delay) > 2s (our timeout).
		fmt.Printf("\nResult: Request Failed (As Expected!)\n")
		fmt.Printf("Time waited: %v\n", duration)
		fmt.Printf("Error Details: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Success!")
}