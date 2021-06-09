package main

import (
	"fmt"
	"time"
)

func DistinctPrimes(n int) int {
	m := make(map[int]bool)

	for n%2 == 0 {
		m[2] = true
		n /= 2
	}

	for i := 3; i*i < n+1; i += 2 {
		for n%i == 0 {
			m[i] = true
			n /= i
		}
	}

	if n > 2 {
		m[n] = true
	}

	return len(m)
}

func main() {
	start := time.Now()
	counter := 0
	n := 2 * 3 * 4 * 7 // first number to have 4 distinct prime factors
	for counter != 4 {
		if DistinctPrimes(n) == 4 {
			counter++
		} else {
			counter = 0
		}
		n++
	}

	fmt.Println(n-4, time.Since(start)) // 134043, 50.6ms
}
