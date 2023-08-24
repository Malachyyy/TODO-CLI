package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func writeToList(filePath string) {
	// Scan user input with spacebars
	addToDo := bufio.NewScanner(os.Stdin)
	// Keep scanning for inputs
	addToDo.Scan()

	input := addToDo.Text()

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Error opening file", err)
	}
	defer f.Close()

	if _, err := f.WriteString(input + "\n"); err != nil {
		fmt.Println("error writing to file", err)
	}

	fmt.Println("Added:", input)

}

func deleteItem() {

}

func readList(filePath string) string {
	todoFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error reading file", err)
	}
	defer todoFile.Close()

	scanner := bufio.NewScanner(todoFile)

	var outputBuilder strings.Builder
	lineNumber := 1

	for scanner.Scan() {
		line := scanner.Text()
		outputBuilder.WriteString(fmt.Sprintf("%d: %s\n", lineNumber, line))
		lineNumber++
	}

	output := outputBuilder.String()

	return output

}
