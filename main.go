package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the path to the JSON file.")
		return
	}
	fileName := os.Args[1]

	jsonData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result := verifyPolicy(jsonData)
	fmt.Println(result)
}
