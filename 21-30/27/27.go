package main

import (
	"fmt"
	"time"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	if n <= 3 {
		return true
	}

	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	start := time.Now()
	maxN, maxA, maxB := 0, 0, 0
	for a := -1000; a <= 1000; a++ {
		for b := -1000; b <= 1000; b++ {
			fmt.Println(a, b)
			streak := 0
			for n := 0; ; n++ {
				if isPrime(n*n + a*n + b) {
					streak++
				} else {
					break
				}
			}
			if streak > maxN {
				maxN = streak
				maxA = a
				maxB = b
			}
		}
	}
	fmt.Println(maxA*maxB, time.Since(start)) // -59231, 37.806583266s
	// Note: slow, maybe optimize?
}
