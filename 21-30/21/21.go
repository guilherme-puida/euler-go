package main

import (
	"fmt"
	"math"
)

func getSumDivisors(a int) int {
	sum := 0
	for i := 1; i < int(math.Sqrt(float64(a)))+1; i++ {
		if a%i == 0 {
			sum += i + a/i
		}
	}
	return sum - a
}

func main() {
	m := make(map[int]bool)

	for a := 1; a < 10000; a++ {
		_, ok := m[a]

		if !ok {
			da := getSumDivisors(a)
			db := getSumDivisors(da)

			if db == a && a != da {
				m[a] = true
				m[da] = true
			}
		}
	}

	sum := 0

	for key := range m {
		sum += key
	}

	fmt.Println(sum) // 31626
}
