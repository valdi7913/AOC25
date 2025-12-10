package main

import (
	s "strings"
)

type Cell struct {
	char rune
	count int
}

type Row struct {
	cells []Cell
}

func FormatInput(input []byte) ([]Row, error) {
	var rows []Row = make([]Row, 0)
	for _, line := range s.Split(string(input),"\n") {
		if line == "" {break}
		cells := make([]Cell, 0)
		for _, r := range []rune(line) {
			cells = append(cells, Cell{char: r, count: 0})
		}
		rows = append(rows, Row{cells: cells})
	}
	return rows, nil
}

func SolvePart1(rows []Row) int{
	splitCount := 0
	for y := range (len(rows) - 1) {
		for x := range (len(rows[y].cells)){
			switch rows[y].cells[x].char {
				case 'S':
					rows[y+1].cells[x].char = '|'
				case '|':
					below := rows[y+1].cells[x].char
					if below == '^' {
						rows[y+1].cells[x+1].char = '|'
						rows[y+1].cells[x-1].char = '|'
						splitCount++
					} else {
						rows[y+1].cells[x].char = '|'
					}
			}
		}
	}
	return splitCount
}

func SolvePart2(rows []Row) int {
	for y := range (len(rows) - 1) {
		for x := range (len(rows[y].cells)){
			switch rows[y].cells[x].char {
				case 'S':
					rows[y+1].cells[x].char = '|'
					rows[y+1].cells[x].count = 1
				case '|':
					below := rows[y+1].cells[x].char
					if below == '^' {
						rows[y+1].cells[x+1].char = '|'
						rows[y+1].cells[x-1].char = '|'
						rows[y+1].cells[x-1].count += rows[y].cells[x].count
						rows[y+1].cells[x+1].count += rows[y].cells[x].count
					} else {
						rows[y+1].cells[x].char = '|'
						rows[y+1].cells[x].count += rows[y].cells[x].count
					}
			}

		}
	}
	sumCount := 0
	for _, cell := range rows[len(rows) - 1 ].cells {
		sumCount += cell.count
	}

	return sumCount
}
