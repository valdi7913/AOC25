package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	file := os.Args[1]

	raw_data, err := os.ReadFile(file)

	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	start := time.Now()
	ranges, ids, err := FormatInput(raw_data)
	elapsed := time.Since(start)

	if err != nil {
		fmt.Println("Error parsing input:", err)
		return
	}

	fmt.Printf("Input parsed in %s\n\n", elapsed)

	start = time.Now()
	solution := SolvePart1(ranges, ids)
	elapsed = time.Since(start)

	fmt.Printf("Solution 1:\n\tFound in %s\n\tSolution: %d\n\n", elapsed, solution)

	start = time.Now()

	solution = SolvePart2(ranges, ids)
	elapsed = time.Since(start)

	fmt.Printf("Solution 2:\n\tFound in %s\n\tSolution: %d\n", elapsed, solution)
}
