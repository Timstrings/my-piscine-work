package main

import (
	"fmt"
	"os"
)

func main() {
	// Define the valid strings
	targets := map[string]bool{
		"01":        true,
		"galaxy":    true,
		"galaxy 01": true,
	}

	// Loop through all command-line arguments (excluding the program name)
	for _, arg := range os.Args[1:] {
		if targets[arg] {
			fmt.Println("Alert!!!")
			return
		}
	}
	// If no match, print nothing
}
