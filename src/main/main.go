package main

import "fmt"
import "helper"

//list test items
const (
	ItemUnknown = iota
	ItemBasics
	ItemFunctions
	ItemArrays
	ItemMaps
)

var trace = helper.Trace
var un = helper.Un
var where = helper.Where

func main() {
	item := ItemBasics
	fmt.Printf("current test item: %v\n", item)

	switch item {
	case ItemBasics:
		runTestBasics()
	case ItemFunctions:
		runTestFunctions()
	case ItemArrays:
		runTestArrays()
	case ItemMaps:
		runTestMaps()
	case ItemUnknown:
		fallthrough
	default:
		fmt.Println("an unknown item to test")
	}
}
