package main

import (
	"fmt"
	"math"
	"time"
)

func countDivisors(n int) int {
	var count int = 0

	for i := 1; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			count += 2
		}
	}
	return count
}

func main() {
	start := time.Now()
	var i int = 1
	for {
		var n int = 0
		for j := 1; j <= i; j++ {
			n += j
		}
		if countDivisors(n) > 500 {
			fmt.Println(n, time.Since(start)) // 76576500, 426.895687ms
			break
		}
		i++
	}
}
