package main

import (
	"fmt"
	"math"
	"strconv"
	"sort"
	s "strings"
)

type item struct {
	x int
	y int
	z int
}

func FormatInput(input []byte) ([]item, error) {
	var items []item = make([]item, 0)
	for i, line := range s.Split(string(input),"\n") {
		if line == "" {break}
		coords := s.Split(line, ",")
		x, xe := strconv.Atoi(coords[0])
		y, ye := strconv.Atoi(coords[1])
		z, ze := strconv.Atoi(coords[2])
		if(xe != nil ||ye != nil ||ze != nil) {
			return items, fmt.Errorf("Error parsing line %d: %s", i, line)
		}
		items = append(items,item{x:x,y:y,z:z})
	}
	return items, nil
}

type Connection struct {
	from int
	to int
	distance float64
}

func SolvePart1(items []item) int{
	fmt.Println("Starting UF with", len(items), "items")
	union := NewUnion(len(items))

	connections := make([]Connection, 0)
	//Calculate distances
	for i, item := range items {
		for j := i; j < len(items); j++ {
			if i == j {continue}
			distance := distance(item, items[j])
			connections = append(connections, Connection{from: i, to: j, distance:distance})
		}
	}

	//Sort distances
	sort.Slice(connections, func(i, j int) bool {
		return connections[i].distance < connections[j].distance
	})

	connectionCount := 10
	for i:= range connectionCount {
		connection := connections[i]
		if !union.Connected(connection.to, connection.from) {
			union.Union(connection.to, connection.from)
		}
	}

	frequencyMap := map[int]int{}
	for i := 0; i < len(union.id); i++ {
		root := union.Find(i)
		value, exists := frequencyMap[root]
		if !exists {
			frequencyMap[root] = 1
		} else {
			frequencyMap[root] = value + 1
		}
	}

	keys := make([]int, 0, len(frequencyMap))

	for key := range frequencyMap {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return frequencyMap[keys[i]] < frequencyMap[keys[j]]
	})

	n := len(frequencyMap)
	ans := frequencyMap[keys[n-1]] * frequencyMap[keys[n-2]] * frequencyMap[keys[n-3]]
	return ans
}

func SolvePart2(items []item) int {
	fmt.Println("Starting UF with", len(items), "items")
	union := NewUnion(len(items))

	connections := make([]Connection, 0)
	for i, item := range items {
		for j := i; j < len(items); j++ {
			if i == j {continue}
			distance := distance(item, items[j])
			connections = append(connections, Connection{from: i, to: j, distance: distance})
		}
	}

	//Sort distances
	sort.Slice(connections, func(i, j int) bool {
		return connections[i].distance < connections[j].distance
	})

	var lastConnection Connection;
	for _, connection := range connections {

		if !union.Connected(connection.to, connection.from) {
			countBeforeUnion:= union.Count()
			union.Union(connection.to, connection.from)
			countAfterUnion:= union.Count()
			if(countBeforeUnion > 1 && countAfterUnion == 1) {
				lastConnection = connection
			}
		}
	}

	ans:=items[lastConnection.from].x * items[lastConnection.to].x
	return ans
}

func distance(a, b item) float64 {
	xdiff:= float64(b.x - a.x)
	ydiff:= float64(b.y - a.y)
	zdiff:= float64(b.z - a.z)
	return math.Sqrt(xdiff*xdiff + ydiff*ydiff + zdiff*zdiff)
}
