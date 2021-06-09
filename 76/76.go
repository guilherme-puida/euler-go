package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var t [101]int
	for i := range t {
		t[i] = 0
	}

	t[0] = 1

	for i := 1; i < 100; i++ {
		for j := i; j < 101; j++ {
			t[j] += t[j-i]
		}
	}

	fmt.Println(t[100], time.Since(start)) // 190569291, 3.706µs
}
