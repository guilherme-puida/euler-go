package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var count int = 0

	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			count += i
		}
	}

	fmt.Println(count, time.Since(start)) // 233168, 1.427µs
}
