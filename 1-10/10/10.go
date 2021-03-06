package main

import (
	"fmt"
	"time"
)

func Erastothenes(n int) []int {
	a := make([]bool, n)
	for i := range a {
		a[i] = true
	}

	for i := 2; i*i < n; i++ {
		if a[i] {
			for j := i * i; j < n; j += i {
				a[j] = false
			}
		}
	}

	var result []int
	for i := range a {
		if a[i] && i > 1 {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	start := time.Now()
	primes := Erastothenes(2000000)
	sum := 0

	for _, e := range primes {
		sum += e
	}

	fmt.Println(sum, time.Since(start)) // 142913828922, 8.156541ms
}
