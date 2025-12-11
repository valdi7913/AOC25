package main

import (
	"fmt"
	"time"
	"os"
)

func main() {
	file := os.Args[1]

	raw_data, err := os.ReadFile(file)

	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	start:= time.Now()
	data, err := FormatInput(raw_data)
	elapsed := time.Since(start)

	if(err != nil) {
		fmt.Println("Error parsing input:", err)
		return
	}

	fmt.Printf("Input parsed in %s\n\n", elapsed)

	copyOfData := make([]item, len(data))
	copy(copyOfData, data)

	start = time.Now()
	solution := SolvePart1(copyOfData)
	elapsed = time.Since(start)

	fmt.Printf("Solution 1:\n\tFound in %s\n\tSolution: %d\n\n", elapsed, solution)

	start = time.Now()

	solution = SolvePart2(data)
	elapsed = time.Since(start)

	fmt.Printf("Solution 2:\n\tFound in %s\n\tSolution: %d\n", elapsed, solution)
}
