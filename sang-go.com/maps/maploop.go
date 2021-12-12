package main

import "fmt"

func main() {

	testMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	for mapKey, mapValye := range testMap {
		fmt.Printf("Key %s, Value %d", mapKey, mapValye)
	}

}
