package main

import (
	"io"
	"os"
)

func main() {
	args := os.Args[1:]

	// No arguments → read from stdin
	if len(args) == 0 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	// Loop through each argument (file)
	for _, file := range args {
		data, err := os.ReadFile(file)
		if err != nil {
			os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
			os.Exit(1) // exit immediately on first error
		}
		os.Stdout.Write(data)
	}
}
