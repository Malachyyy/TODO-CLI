package main

import (
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
		fmt.Println(readFile(filePath))
	case "T":
		test(filePath)
	default:
		fmt.Println("Insert the correct option ")

	}
}

func Delete(filePath string) {
	outputList := numberFile(filePath)

	fmt.Println("Which line do you wanna delete?")
	fmt.Println(outputList)

	var lineToDelete int
	fmt.Scan(&lineToDelete)

	newContent := deleteItem(outputList, lineToDelete)
	fmt.Println(newContent)
	fmt.Println("Deleted line:", lineToDelete)
}

func Add(filePath string) {
	fmt.Println("Add an item")

	writeChanges(filePath)

}

func test(filePath string) {
	fmt.Println(readFile(filePath))

}
