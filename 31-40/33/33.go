package main

import (
	"fmt"
	"time"
)

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func main() {
	start := time.Now()
	d, n := 1, 1

	for i := 1; i < 10; i++ {
		for j := 1; j < i; j++ {
			for k := 1; k < j; k++ {
				if (k*10+i)*j == k*(i*10+j) {
					d *= j
					n *= k
				}
			}
		}
	}
	d /= GCD(d, n)
	fmt.Println(d, time.Since(start)) // 100, 542ns
}
