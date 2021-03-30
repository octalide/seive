package main

import (
	"fmt"
	"time"
)

var (
	limit = 100000000

	primes []bool
)

func create(size int) []bool {
	slice := make([]bool, limit)

	for i := 0; i < size; i++ {
		slice[i] = true
	}

	return slice
}

func seive(lower, upper int) []bool {
	offset := upper - lower

	p := create(offset)

	for num := lower; num < upper; num++ {
		// ignores 0, 1, and 2
		if num <= 2 {
			continue
		}

		// removes even numbers
		if num%2 == 0 {
			p[num] = false
			continue
		}

		// removes multiples of the current number
		for mult := num * 2; mult < upper; mult += num {
			p[mult] = false
		}
	}

	return p
}

func main() {
	fmt.Println("limit:      ", limit)

	start := time.Now()
	primes = seive(0, limit)
	duration := time.Since(start)

	// create an integer list with all the calculted primes
	indices := make([]int, 0)
	for i, p := range primes {
		// ignore 0 and 1
		if i <= 1 {
			continue
		}

		if p {
			indices = append(indices, i)
		}
	}

	fmt.Println("prime count:", len(indices))
	fmt.Println("time:       ", duration)
}
