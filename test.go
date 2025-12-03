package main

import (
	"fmt"
)

func main () {
	for i:=0; i < 10; i++ {

		fmt.Println("%d, %v", i, Factors(i))
	}
}

func Factors(number int) []int {
	primes := make([]int, 0)

	if(number > 0) {
		primes = append(primes, 1)
	}

	for i:= 2; i * i <= number ; i++ {
		if number % i == 0 {
			primes = append(primes, i)
		}
	}

	if(number > 1) {
		primes = append(primes, number)
	}

	return primes
}
