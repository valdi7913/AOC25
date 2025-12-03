package main

import (
	"fmt"
	s "strings"
)

type item struct {

}

func FormatInput(input []byte) ([]item, error) {
	var items []item = make([]item, 0)
	for i, line := range s.Split(string(input),",") {
		fmt.Println(i, line)
	}
	return items, nil
}

func SolvePart1(items []item) int{
	return 0
}

func SolvePart2(items []item) int {
	return 0
}
