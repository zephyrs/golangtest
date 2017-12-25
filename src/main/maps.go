package main

import (
	"fmt"
)

func runTestMaps() {
	defer un(trace("runTestMaps"))

	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 2)
		items[i][1] = 2
		items[i][2] = 4
	}
	fmt.Printf("Version A: Value of items: %v\n", items)

	mapCreated := map[string]float32{"0": 4.5, "1": 89.1, "2": 99.2}
	for _, v := range mapCreated {
		fmt.Println("Map item: ", v)
	}

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapCreated["key3"] = 3.1
	mapCreated["key4"] = 2.1

	for _, v := range mapCreated {
		fmt.Println("Map item: ", v)
	}

	v := mapCreated["key1"]
	delete(mapCreated, "key1")
	fmt.Println(v)

	capitals := map[string]string{"France": "Paris", "Japan": "Tokyo", "Italy": "Rome"}
	for key, val := range capitals {
		fmt.Println("Map item: Capital of", key, "is", val)
	}
}
