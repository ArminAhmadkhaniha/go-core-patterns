package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"io"

)
func main() {
	url := "https://github.com/ArminAhmadkhaniha"

	fmt.Println("--- 1. The http.Get Method ---")

	// METHOD: http.Get(url)
	// WHAT IT DOES: It acts like opening a browser tab. It sends a request and waits.
	// RETURNS: Two things:
	//   1. 'resp': The envelope containing the HTML, Status Code, and Headers.
	//   2. 'err': If the INTERNET failed (e.g., WiFi off, invalid domain).
	resp, err := http.Get(url)
	if err != nil {
		// We use 'os.Exit(1)' to stop the program immediately with an error code.
		fmt.Println("Could not connect",err)
		os.Exit(1)

	}
	// --- Memory Management---
	// METHOD: resp.Body.Close()
	// WHY: When you download a webpage, Go keeps a "stream" open (like a water tap).
	// If you don't close it, you leak memory. 'defer' ensures it closes when main() finishes.
	defer resp.Body.Close()

	fmt.Println("--- 2. Checking the Status Code ---")
	// FIELD: resp.StatusCode
	// WHAT IT DOES: Tells you the result (200 = OK, 404 = Not Found, 500 = Server Error).
	fmt.Println("Status Code:",resp.StatusCode)

	if resp.StatusCode == 200 {
		fmt.Println("Result: Success! Your profile is live.")
	} else if resp.StatusCode == 404 {
		fmt.Println("Result: Failed. Profile not found.")
	}

	fmt.Println("--- 3. Reading the Body (The Content) ---")
	// METHOD: io.ReadAll(resp.Body)
	// WHAT IT DOES: It reads the actual HTML text inside the envelope.
	// NOTE: We only read the first 100 characters to keep the terminal clean.

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Convert bytes to string and print the first 100 chars
	content := string(bodyBytes)
	fmt.Printf("Preview of Page Content:\n%s...\n", content[:100])
}


	



	

