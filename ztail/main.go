package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		return
	}
	if os.Args[1] != "-c" {
		return
	}
	n := 0
	for _, ch := range os.Args[2] {
		if ch < '0' || ch > '9' {
			return
		}
		n = n*10 + int(ch-'0')
	}
	files := os.Args[3:]
	exitCode := 0
	printedAny := false

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			exitCode = 1
			printedAny = true
			continue
		}
		if printedAny {
			fmt.Println()
		}
		if len(files) > 1 {
			fmt.Printf("==> %s <==\n", file)
		}
		if len(data) > n {
			fmt.Print(string(data[len(data)-n:]))
		} else {
			fmt.Print(string(data))
		}
		printedAny = true
	}
	if exitCode != 0 {
		os.Exit(1)
	}
}
