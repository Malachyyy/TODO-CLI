package main

import (
	"bufio"
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println("TODO-LIST-CLI (A)dd (D)elete (E)dit (R)ead (S)top ")

	filePath := "text.txt"

	var input string

	fmt.Scan(&input)

	switch input {
	case "A":
		Add(filePath)
	case "E":
		fmt.Println("You are editing")
	case "D":
		Delete(filePath)
	case "S":
		fmt.Println("Stopping program...")
		os.Exit(0)
	case "R":
		fmt.Println(readList(filePath))
	case "T":
		test(filePath)
	default:
		fmt.Println("Insert the correct option ")

	}
}

func Delete(filePath string) {
	fmt.Println("What line do you wanna delete?")

	readList(filePath)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "unwanted" {
			fmt.Println(line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error deleting line", err)
	}
}

func Add(filePath string) {
	fmt.Println("Add an item")

	writeToList(filePath)

}

func test(filePath string) {
	fmt.Println(readList(filePath))

}
