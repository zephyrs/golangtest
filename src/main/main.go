package main

import "fmt"
import "helper"

//list test items
const (
	Unknown = iota
	Basics
	Functions
)

var trace = helper.Trace
var un = helper.Un

func main() {
	item := Functions
	fmt.Printf("current test item: %v\n", item)

	switch item {
	case Basics:
		runTestBasics()
	case Functions:
		runTestFunctions()
	default:
		fmt.Println("an unknown item to test")
	}
}
