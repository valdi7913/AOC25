package main

import (
	"fmt"
	"strconv"
	s "strings"
	"sort"
)

type Point struct {
	x int
	y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func FormatInput(input []byte) ([]Point, error) {
	var points []Point = make([]Point, 0)
	for i, line := range s.Split(string(input),"\n") {
		if line == "" {break}
		arr := s.Split(line, ",")

		if (len(arr) != 2) {
			return points, fmt.Errorf("Error parsing line %d: %s", i, line)
		}

		x, xerror := strconv.Atoi(arr[0])
		y, yerror := strconv.Atoi(arr[1])

		if xerror != nil || yerror != nil {
			return points, fmt.Errorf("Error parsing line %d: %s, x error %s, y error %s", i, line, xerror, yerror)
		}

		points = append(points, Point{x:x, y:y})
	}
	return points, nil
}

type Rectangle struct {
	ul Point
	lr Point
}

func NewRectangle(a, b Point) Rectangle{
	minX := min(a.x, b.x)
	maxX := max(a.x, b.x)
	minY := min(a.y, b.y)
	maxY := max(a.y, b.y)

	return Rectangle{ul: Point{x: minX, y: minY}, lr: Point{x: maxX, y: maxY}}
}

func (r *Rectangle) Area() int {
	w := r.lr.x - r.ul.x + 1
	h := r.lr.y - r.ul.y + 1

	if w < 0 {w *= -1}
	if h < 0 {w *= -1}
	return w * h
}

func (r Rectangle) String() string {
	return fmt.Sprintf("UL:%s LR:%s Area:%d\n", r.ul.String(), r.lr.String(), r.Area())
}

func SolvePart1(points []Point) int{

	rectangles := make([]Rectangle, 0)

	for i, ul := range points {
		for j, lr := range points {
			if i == j {continue}
			rectangles = append(rectangles, NewRectangle(ul,lr))
		}
	}

	sort.Slice(rectangles, func(i, j int) bool {
		return rectangles[i].Area() > rectangles[j].Area()
	})

	// fmt.Println("highest area", rectangles[0])
	// draw(points, rectangles[0])
	return rectangles[0].Area()
}

func SolvePart2(points []Point) int {
	return 0
}

func draw(tiles []Point, rectangle Rectangle) {
	fmt.Println("Drawing rectangle:", rectangle)
	// Source - https://stackoverflow.com/a
	// Posted by nmichaels, modified by community. See post 'Timeline' for change history
	// Retrieved 2025-12-11, License - CC BY-SA 3.0
	maxInt := int(^uint(0) >> 1)

	maxX := -1
	maxY := -1
	minX := maxInt
	minY := maxInt

	for _, point := range tiles {
		if point.x > maxX {
			maxX = point.x
		}
		if point.x < minX {
			minX = point.x
		}
		if point.y > maxY {
			maxY = point.y
		}
		if point.y < minY {
			minY = point.y
		}
	}

	UL := Point{x: minX, y: minY}
	LR := Point{x: maxX, y: maxY}

	rectangle.ul.x -= UL.x
	rectangle.ul.y -= UL.y

	n := LR.y - UL.y + 1
	m := LR.x - UL.x + 1
	grid := make([][]rune, n)
	for y := range n {
		grid[y] = make([]rune, m)
		for x := range m {
			if rectangle.ul.x <= x && x <= rectangle.lr.x &&
				 rectangle.ul.y <= y && y <= rectangle.lr.y {
				grid[y][x] = 'O'
			} else {
				grid[y][x] = '.'
			}
		}
	}

	for i := range tiles {
		tiles[i].x -= UL.x
		tiles[i].y -= UL.y
		grid[tiles[i].y][tiles[i].x] = '#'
	}

	// for range (len(grid[0]) + 2) {
	// 	fmt.Print(".")
	// }
	fmt.Println()
	for _, row := range grid {
		// fmt.Print(".")
		for _, cell := range row {
			fmt.Printf("%s", string(cell))
		}
		// fmt.Print(".")
		fmt.Println()
	}
	// for range (len(grid[0]) + 2) {
	// 	fmt.Print(".")
	// }
	fmt.Println()
}
