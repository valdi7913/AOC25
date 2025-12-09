package main

import (
	"fmt"
	s "strings"
	"regexp"
	"strconv"
)

func FormatInput(input []byte) ([][]int, []string, error) {
	re := regexp.MustCompile(" +")
	var grid [][]int = make([][]int, 0)
	var ops []string = make([]string, 0)
	for line := range s.SplitSeq(string(input),"\n") {
		if line == "" {break}
		line = s.Trim(line, " ")
		line = re.ReplaceAllString(line, ",")
		nums := make([]int, 0)
		for item := range s.SplitSeq(line, ",") {
			if s.Contains(line, "+") || s.Contains(line, "*") {
				ops = append(ops, item)
			} else {
				num, e := strconv.Atoi(item)
				if e != nil {
					return grid, ops, fmt.Errorf("%s", e)
				}
				nums = append(nums, num)
			}
		}
		if !s.Contains(line, "+") && !s.Contains(line, "*") {
			grid = append(grid, nums)
		}
	}

	return grid, ops, nil
}

func Split(r rune) bool {
	return r=='*' || r=='+'
}

func FormatInput2(input []byte) ([][]int, []string, error) {
	var grid [][]int = make([][]int, 0)
	var ops []string = make([]string, 0)
	max := 0
	for line := range s.SplitSeq(string(input),"\n") {
		if max < len(line) {
			max = len(line)
		}
	}

	colWidths := make([]int, 0)
	//Find col widths
	for line := range s.SplitSeq(string(input), "\n") {
		if line == "" {continue}
		if s.Contains(line, "+") || s.Contains(line, "*") {
			curr := 0
			for i := 0; i < max; i++ {
				if i < len(line)-1 && line[i+1] != ' ' {
					colWidths = append(colWidths, curr)
					curr = 0
				} else {
					curr++
				}
				if i == max - 1 {
					colWidths = append(colWidths, max - len(line) + 1)
				}
			}
		}
	}


	lines := s.Split(string(input), "\n")
	//Pad 0s
	for i := 0; i<len(lines); i++ {
		if lines[i] == "" {continue}
		if s.Contains(lines[i], "+") || s.Contains(lines[i], "*") {continue}

	}

	//Pad 0s
	// for line := range s.SplitSeq(string(input),"\n") {
	// 	if line == "" {break}
	// 	// nums := make([]int, 0)
	// 	for i:= 0; i < max; i++ {
	// 		if i >= len(line) {
	// 			fmt.Print(" ")
	// 		} else {
	// 			fmt.Print(string(line[i]))
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	//Parse numbers
	//
	//

	return grid, ops, nil
}

func SolvePart1(nums [][]int, ops []string) int{
	// fmt.Printf("%v", nums)
	// fmt.Println("cols ", len(nums), "row:", len(nums[0]))
	grand_sum:=0
	cols := len(nums[0])
	rows := len(nums)

	for col:= 0; col<cols; col++ {
		sum := 0
		prod := 1
		for row:= 0; row<rows; row++ {
			if(ops[col] == "+") {
				sum += nums[row][col];
				print("+", nums[row][col])
			}
			if(ops[col] == "*") {
				prod *= nums[row][col]
				print("*", nums[row][col])
			}
		}
		println()
		if(ops[col] == "+") {
			grand_sum += sum
			println(col, "+ solution:", sum)
		} else {
			println(col, "* solution:", prod)
			grand_sum += prod
		}
	}
	return grand_sum
}

func SolvePart2(nums [][]int, ops []string) int {
	return 0
}
