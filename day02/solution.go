package main

import (
	"fmt"
	"strconv"
	s "strings"
)

type item struct {
	first int
	last int
}

func FormatInput(input []byte) ([]item, error) {
	var items []item = make([]item, 0)
	for i, line := range s.Split(string(input),",") {
		line = s.ReplaceAll(line, "\n", "")

		arr := s.Split(line, "-")
		if(len(arr) != 2) {
			return items, fmt.Errorf("Error parsing line %d: %s is too short\n", i, line)
		}

		first, fe := strconv.Atoi(arr[0])
		last, le := strconv.Atoi(arr[1])

		if(fe != nil){
			return items, fmt.Errorf("Error parsing first item on line %d: %s, %s\n", i, line, fe)
		}

		if(le != nil) {
			return items, fmt.Errorf("Error parsing last item on line %d: %s, %s\n", i, line, le)
		}

		items = append(items, item{first:first, last:last})
	}
	return items, nil
}

func SolvePart1(items []item) int{
	sum:= 0
	for _, item := range items {
		for i:= item.first; i < item.last; i++ {
			if(!IdIsValid(i)) {
				sum+=i
			}
		}
	}
	return sum
}

func SolvePart2(items []item) int {
	sum:= 0
	for _, item := range items {
		for i:= item.first; i <= item.last; i++ {
			if(IdHasRepeatingPatterns(i)) {
				sum+=i
			}
		}
	}
	return sum
}


func IdIsValid(id int) bool {
	str := strconv.Itoa(id)

	if(len(str) % 2 == 1) {return true}

	half := len(str)/2
	if(str[0:half] == str[half:]) {
		return false
	}

	return true
}

func IdHasRepeatingPatterns(id int) bool {
	str := strconv.Itoa(id)
	factors := Factors(len(str))

	for _, factor := range factors {
		var last string = str[0: factor]
		in:
		for j := factor; j < len(str); j+=factor {
			current := str[j:j+factor]
			if current != last && last != "" {break in}
			last = current
			if j+factor > len(str) - 1 {
				return true
			}
		}
	}
	return false
}

func Factors(number int) []int {
	primes := make([]int, 0)

	for i:= 1; i <= number/2 + 1; i++ { // ekki 100% efficient en gæti ekki verið meira sama
		if number % i == 0 {
			primes = append(primes, i)
		}
	}

	return primes
}
