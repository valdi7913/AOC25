package main

import (
	"fmt"
	s "strings"
)

type Row struct {
	cells []Cell
}

type Cell struct {
	full bool
}

func (r Row) String() string {
	return fmt.Sprintf("Row: %v\n", r.cells)
}

func FormatInput(input []byte) ([]Row, error) {
	var rows []Row = make([]Row, 0)
	for i, line := range s.Split(string(input),"\n") {
		row := Row{cells: make([]Cell, 0)}
		if(len(line) == 0) {break}
		for j, letter := range s.Split(line,"") {
			fmt.Println(i, j, letter)
			row.cells = append(row.cells, Cell{full: letter=="@"})
		}
		rows = append(rows, row)
	}
	return rows, nil
}

type Point struct {
	x int
	y int
}

func SolvePart1(rows []Row) int{
	l := len(rows)
	accessible := 0
	for i, row := range rows {
		for j, cell := range row.cells {
			if !cell.full {continue}
			neighbors := []Point{
				{x:i-1, y:j-1},
				{x:i-1, y:j+0},
				{x:i-1, y:j+1},
				{x:i+0, y:j-1},
				{x:i+0, y:j+0},
				{x:i+0, y:j+1},
				{x:i+1, y:j-1},
				{x:i+1, y:j+0},
				{x:i+1, y:j+1},
			}

			count := 0
			// neighbors:
			for _, n := range neighbors {
				if(n.x >= 0 && n.y >= 0 && n.x < l && n.y < l && rows[n.x].cells[n.y].full) {
					count++;
				}
			}
			fmt.Println("Cell in", i, " col ", j, "has", count, "neighbors" )
			if(count < 4) {
				accessible++
			}
		}
	}
	return accessible
}

func SolvePart2(rows []Row) int {
	return 0
}
