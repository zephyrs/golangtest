package main

import "fmt"

//list test items
const (
	Basics = iota
	Functions
)

func main() {
	item := Functions
	fmt.Printf("current test item: %v", item)

	switch item {
	case Basics:
		runTestBasics()
	case Functions:
		runTestFunctions()
	default:
		fmt.Println("an unknown item to test")
	}
}
