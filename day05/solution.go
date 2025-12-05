package main

import (
	"fmt"
	"strconv"
	s "strings"
)

type Range struct {
	start int
	end   int
}

func FormatInput(input []byte) ([]Range, []int, error) {
	var ranges []Range = make([]Range, 0)
	var ids []int = make([]int, 0)

	ranges_finished := false
	for i, line := range s.Split(string(input), "\n") {
		if line == "" && !ranges_finished {
			ranges_finished = true
			continue
		} else if line == "" {
			break
		}
		if !ranges_finished {
			arr := s.Split(line, "-")
			start, se := strconv.Atoi(arr[0])
			end, ee := strconv.Atoi(arr[1])
			if se != nil || ee != nil {
				return ranges, ids, fmt.Errorf("Unable to parse line %d, %s, %s, %s", i, line, ee, se)
			}
			ranges = append(ranges, Range{start: start, end: end})
		} else {
			id, ie := strconv.Atoi(line)
			if ie != nil {
				return ranges, ids, fmt.Errorf("Unable to parse line %d, %s, %s", i, line, ie)
			}
			ids = append(ids, id)
		}
	}
	return ranges, ids, nil
}

func SolvePart1(ranges []Range, ids []int) int {
	fmt.Println(len(ranges))
	fmt.Println(len(ids))
	freshIngredients := 0
	for _, id := range ids {
	ranges:
		for _, r := range ranges {
			if r.start <= id && id <= r.end {
				freshIngredients++
				break ranges
			}
		}
	}

	return freshIngredients
}

func SolvePart2(items []Range, ids []int) int {
	return 0
}
