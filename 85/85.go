package main

import "fmt"

func countRectangles(n int, m int) int {
	return (m * n * (n + 1) * (m + 1)) / 4
}

func abs(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func main() {
	i := 1
	for {
		for j := 1; j <= i; j++ {
			current := countRectangles(i, j)
			if abs(current, 2000000) < 100 {
				fmt.Println(i, j, current, i*j) // 2772
			}
		}
		i++
	}
}
