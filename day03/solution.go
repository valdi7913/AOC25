package main

import (
	"fmt"
	"strconv"
	s "strings"
)

type Bank struct {
	batteries []Battery
}

func (b Bank) String() string {
	return fmt.Sprintf("Bank: %v\n", b.batteries)
}

type Battery struct {
	value int
}

func FormatInput(input []byte) ([]Bank, error) {
	var banks []Bank = make([]Bank, 0)
	for _, line := range s.Split(string(input),"\n") {
		if len(line) == 0 {break}
		var bank Bank = Bank{batteries: make([]Battery,0)}
		for rune := range s.SplitSeq(line, "") {
			value, err := strconv.Atoi(rune)
			if err != nil {
				return banks, fmt.Errorf("Unable to parse %s in line %s, error", rune, line, err)
			}
			bank.batteries = append(bank.batteries, Battery{value})
		}
		banks = append(banks, bank)
	}
	return banks, nil
}

func SolvePart1(banks []Bank) int{
	sum := 0
	for _, bank := range banks {
		maxLeftIndex:= 0
		maxLeftValue:= bank.batteries[maxLeftIndex].value
		for i := 0; i < len(bank.batteries) - 1; i++ {
			if bank.batteries[i].value > maxLeftValue {
				maxLeftIndex = i;
				maxLeftValue = bank.batteries[i].value
			}
		}

		maxRightIndex:= maxLeftIndex + 1
		maxRightValue:= bank.batteries[maxRightIndex].value
		for i := maxLeftIndex+1; i < len(bank.batteries); i++ {
			if bank.batteries[i].value > maxRightValue {
				maxRightIndex = i
				maxRightValue = bank.batteries[i].value
			}
		}

		// fmt.Println(10 * maxLeftValue + maxRightValue, "leftIndex", maxLeftIndex, "rightIndex", maxRightIndex, "batteries", bank)
		sum+= 10 * maxLeftValue + maxRightValue
	}
	return sum
}

func SolvePart2(banks []Bank) int {
	sum := 0
	max := 12
	for _, bank := range banks {
		maxJolt:=0
		maxLeftIndex:= 0
		for j := max; j > 0; j-- {
			maxLeftValue:= bank.batteries[maxLeftIndex].value
			for i := maxLeftIndex; i < len(bank.batteries) - j +1; i++ {
				if bank.batteries[i].value > maxLeftValue {
					maxLeftIndex = i;
					maxLeftValue = bank.batteries[i].value
				}
			}
			maxJolt+= intPow(10,j-1) * maxLeftValue
			maxLeftIndex += 1
		}
	 	sum += maxJolt
	}
	return sum
}

func intPow(base, exp int) int {
	result := 1
	for exp > 0 {
		result *= base
		exp--
	}
	return result
}
