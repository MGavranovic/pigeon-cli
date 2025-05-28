package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World!\nWelcome to pigeon-cli!")

	dir, err := os.ReadDir("./")
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	for _, file := range dir {
		fmt.Printf("File: %s\n", file.Name())
	}

	inFileSystem := true

	for inFileSystem {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error accessing path: %s", err)
		}
		fmt.Printf("%s:", wd)

		var input string
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		}
	}
}
