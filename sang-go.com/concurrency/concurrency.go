package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	go func() { // 1/ our first go routine fire
		defer waitGrp.Done()        // 7
		time.Sleep(5 * time.Second) // 3
		fmt.Println("Hello")        // 6
	}()

	go func() { // 2/ second go routine created but get any time on thread until first sleep
		defer waitGrp.Done()       //5
		fmt.Println("Pluralsight") // 4/ it pump out pluralsight and reported to the waitGrp it was done and back to the first
	}()

	waitGrp.Wait() // 8

}
