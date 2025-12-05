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
	s := ""
	for _, cell := range r.cells {
		if cell.full {
			s += fmt.Sprint("@")
		} else {
			s += fmt.Sprint(".")
		}
	}
	s += fmt.Sprintln()
	return s
}

func FormatInput(input []byte) ([]Row, error) {
	var rows []Row = make([]Row, 0)
	for line := range s.SplitSeq(string(input), "\n") {
		row := Row{cells: make([]Cell, 0)}
		if len(line) == 0 {
			break
		}
		for letter := range s.SplitSeq(line, "") {
			row.cells = append(row.cells, Cell{full: letter == "@"})
		}
		rows = append(rows, row)
	}
	return rows, nil
}

type Point struct {
	x int
	y int
}

func (p Point) String() string {
	return fmt.Sprintf("x: %d, y: %d\n", p.x, p.y)
}

func SolvePart1(rows []Row) int {
	fmt.Print(findRemoveable(rows))
	return len(findRemoveable(rows))
}

func SolvePart2(rows []Row) int {
	removable := findRemoveable(rows)
	totalRemoved := len(removable)
	for len(removable) > 0 {

		for _, point := range removable {
			rows[point.x].cells[point.y].full = false
		}

		removable = findRemoveable(rows)
		totalRemoved += len(removable)
	}
	fmt.Println(rows)
	return totalRemoved
}

func findRemoveable(rows []Row) []Point {
	l := len(rows)
	removable := []Point{}
	for i, row := range rows {
		for j, cell := range row.cells {
			if !cell.full {
				continue
			}
			neighbors := []Point{
				{x: i - 1, y: j - 1},
				{x: i - 1, y: j + 0},
				{x: i - 1, y: j + 1},
				{x: i + 0, y: j - 1},
				{x: i + 0, y: j + 1},
				{x: i + 1, y: j - 1},
				{x: i + 1, y: j + 0},
				{x: i + 1, y: j + 1},
			}

			// neighbors:
			count := 0
			for _, n := range neighbors {
				if n.x >= 0 && n.y >= 0 && n.x < l && n.y < l && rows[n.x].cells[n.y].full {
					count++
				}
			}
			if count < 4 {
				removable = append(removable, Point{x: i, y: j})
			}
		}
	}
	return removable
}
