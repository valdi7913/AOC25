package main

import (
	"fmt"
	s "strings"
	"strconv"
)


type operation struct {
	is_left bool
	value int
}

func FormatInput(input []byte) ([]operation, error){
	var operations []operation = make([]operation, 0)

	for i, line := range s.Split(string(input), "\n") {
		if len(line) <= 1 {
			if line == "" {
				continue
			}
			return operations, fmt.Errorf("Error parsing line %d: %s is too short\n", i, line)
		}

		value, err := strconv.Atoi(line[1:])
		if err != nil {
			return operations, fmt.Errorf("Error parsing line %d:instruction is not valid %s\n", i, line)
		}

		is_left := false
		switch line[0] {
			case 'L':
				is_left = true
			case 'R':
				is_left = false
			default:
				return operations, fmt.Errorf("Error parsing line %d: invalid direction %s\n", i, line)
		}

		operations = append(operations, operation{is_left: is_left, value: value})
	}
	return operations, nil
}

func SolvePart1(ops []operation) int{
	var dial int = 50
	var password int = 0
	for _, op := range ops {
		if op.is_left {
			dial -= op.value
		} else {
			dial += op.value
		}

		dial %= 100

		if dial == 0 {
			password++
		};
	}
	return password
}

type tuple struct {
	A int
	B int
}

func SolvePart2(ops []operation) int {
	dial := 50
	oldDial := 50
	password := 0
	for _, op := range ops {
		pdiff := 0
		ddiff := 0

		if op.is_left {
			ddiff = (dial - op.value)
		} else {
			ddiff = (dial + op.value)
		}

		dial = ddiff % 100

		if dial < 0 { dial += 100 }

		if dial == 0 && oldDial != 0 {
			pdiff += 1
		}

		if op.is_left && oldDial - op.value%100 < 0 && oldDial != 0 {
			pdiff += 1
		} else if !op.is_left && oldDial + op.value%100 > 100 && oldDial != 0{
			pdiff += 1
		}
		pdiff += op.value / 100

		oldDial = dial
		password += pdiff
		// fmt.Printf("#%d, left %t, op %d, %d\n", i, op.is_left, op.value, dial)

	}

	return password;
}
