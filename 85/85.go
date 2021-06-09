package main

import (
	"fmt"
	"time"
)

func countRectangles(n, m int) int {
	return (m * n * (n + 1) * (m + 1)) / 4
}

func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func main() {
	start := time.Now()
	i := 1
out:
	for {
		for j := 1; j <= i; j++ {
			current := countRectangles(i, j)
			if abs(current, 2000000) < 100 {
				fmt.Println(i*j, time.Since(start)) // 2772, 4.455µs
				break out
			}
		}
		i++
	}
}
