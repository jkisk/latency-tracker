package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	// create a data file
	f, err := os.Create("kay.txt")
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	// check for error below while closing file at end?
	defer f.Close()
	for i := 0; i < 9500; i++ {
		// min represents our fastest server response time
		min := 100
		// max represents a slow but still acceptable response time
		max := 300
		// generate a random int in the above range and write to file
		dataPoint := rand.Intn(max-min+2) + min
		_, err = f.WriteString(fmt.Sprintf("%d\n", dataPoint))
		if err != nil {
			fmt.Printf("error writing int: %v", err)
		}
	}
	// add some outlier data representing slow responses
	for i := 0; i < 500; i++ {
		min := 301
		max := 30000
		// generate a random int in the above range and write to file
		dataPoint := rand.Intn(max-min+1) + min
		_, err = f.WriteString(fmt.Sprintf("%d\n", dataPoint))
		if err != nil {
			fmt.Printf("error writing int: %v", err)
		}
	}
}
