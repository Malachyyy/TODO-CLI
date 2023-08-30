package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func writeChanges(filePath string) {
	// Add items to the text file
	addToDo := bufio.NewScanner(os.Stdin)

	addToDo.Scan()

	input := addToDo.Text()

	write, err := readFileWhileWriting(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Error opening file", err)
	}
	defer write.Close()

	if _, err := write.WriteString(input + "\n"); err != nil {
		fmt.Println("error writing to file", err)
	}

	fmt.Println("Added:", input)

}

func deleteItem(outputList string, lineToDelete int) string {
	lines := strings.Split(outputList, "\n")

	if lineToDelete < 1 || lineToDelete > len(lines) {
		return "Invalid line number"
	}

	// Remove the line at the specified line number.
	lines = append(lines[:lineToDelete-1], lines[lineToDelete:]...)

	// Rebuild the content without the deleted line.
	newContent := strings.Join(lines, "\n")

	return newContent

}

func numberFile(filePath string) string {
	// Returns the text file with linenumbers
	todoFile, err := readFile(filePath)
	if err != nil {
		fmt.Println("Error reading file from numberFile", err)
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

func readFileWhileWriting(filePath string, flags int, perm os.FileMode) (*os.File, error) {
	// Read text file with flags and permissions to write to
	todoFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return todoFile, err

}

func readFile(filePath string) (*os.File, error) {
	todoFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return todoFile, err

}
