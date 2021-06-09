package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	a, b, c := 0, 0, 0

out:
	for a = 1; a < 334; a++ {
		for b = a; b < 500; b++ {
			c = 1000 - a - b

			if a*a+b*b == c*c {
				break out
			}
		}
	}

	fmt.Println(a*b*c, time.Since(start)) // 31875000, 81.213µs
}
