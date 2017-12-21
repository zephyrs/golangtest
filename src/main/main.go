package main

import "fmt"
import "helper"

//list test items
const (
	ItemUnknown = iota
	ItemBasics
	ItemFunctions
)

var trace = helper.Trace
var un = helper.Un
var where = helper.Where

func main() {
	item := ItemFunctions
	fmt.Printf("current test item: %v\n", item)

	switch item {
	case ItemBasics:
		runTestBasics()
	case ItemFunctions:
		runTestFunctions()
	case ItemUnknown:
		fallthrough
	default:
		fmt.Println("an unknown item to test")
	}
}
