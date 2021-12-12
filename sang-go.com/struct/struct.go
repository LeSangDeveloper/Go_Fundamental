package main

import "fmt"

func main() {

	type courseMeta struct {
		author string
		level  string
		rating float64
	}

	gettingStratedWithK8s := courseMeta{
		author: "legel",
		level:  "beginner",
		rating: 5,
	}

	fmt.Printf("%v", gettingStratedWithK8s)
}
