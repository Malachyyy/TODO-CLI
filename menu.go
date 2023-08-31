package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func showMenu() {
	fmt.Println("TODO-LIST-CLI")
	fmt.Println("(A)dd Task")
	fmt.Println("(E)dit Task")
	fmt.Println("(D)elete Task")
	fmt.Println("(R)ead Tasks")
	fmt.Println("(S)top")

	filePath := "text.txt"

	var input string
	fmt.Scan(&input)

	input = strings.ToUpper(input)

	testDeleteItem(filePath, 1)

	/*switch input {
	case "A":
		Add(filePath)
	case "E":
		fmt.Println("You are editing")
	case "D":
		Delete(filePath)
	case "S":
		fmt.Println("Exited program.")
		os.Exit(0)
	case "R":
		fmt.Println(readFile(filePath))
	default:
		fmt.Println("Insert the correct option ")

	}*/
}

func testDeleteItem(filePath string, lineToDelete int) {
	todoFile, err := readFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer todoFile.Close()

	lineToDelete = 3

	var buf bytes.Buffer
	io.Copy(&buf, todoFile)
	asString := string(buf.Bytes())

	output := deleteItem(asString, lineToDelete)

	fmt.Println(output)
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

	writeItemsToFile(filePath)

}
