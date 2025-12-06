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

func SolvePart2(ranges []Range, ids []int) int {
	oldRanges := []Range{ranges[0]}
	count := ranges[0].end - ranges[0].start + 1
	for n := 0; n < len(ranges); n++ {
		new := ranges[n]
		oldRanges:
		for o := 0; o < len(oldRanges); o++ {
			old := oldRanges[o]
			if new.end < new.start {break oldRanges}
			if old.start <= new.start && new.start <= old.end && old.end < new.end { // LEFT
				// fmt.Printf("new start %d end %d\n old start: %d end %d\n moving LEFT\n", new.start, new.end, old.start, old.end)
				new.start = old.end + 1
			} else if new.start < old.start && old.start <= new.end && new.end <= old.end { // RIGHT
				// fmt.Printf("new start %d end %d\n old start: %d end %d\n moving RIGHT\n", new.start, new.end, old.start, old.end)
				new.end = old.start - 1
			} else if old.start <= new.start && new.end <= old.end { // OUTER
				// fmt.Printf("new start %d end %d\n old start: %d end %d\n range is contained\n", new.start, new.end, old.start, old.end)
				new.start = new.end + 1
				break oldRanges
			} else if new.start < old.start && old.end < new.end { // INNER
				// fmt.Printf("new start: %d end %d\nold start: %d end %d\nrange is split into\n\t1: %d-%d and \n\t2: %d-%d\n\n", new.start, new.end, old.start, old.end, new.start, old.start - 1, old.end + 1, new.end)
				//Splits new into 2, add right one to list, continue with left one
				ranges = append(ranges, Range{start: old.end + 1, end: new.end})
				ranges = append(ranges, Range{start: new.start, end: old.start -1})
				new.start = new.end + 1
				break oldRanges
			}
		}
		if new.end >= new.start {
			count += new.end - new.start + 1
			oldRanges = append(oldRanges, Range{start: new.start, end: new.end})
		}
	}

	return count
}
