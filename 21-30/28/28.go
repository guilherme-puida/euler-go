package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	f := make(map[int]int)
	f[0] = 1

	for n := 1; n <= 500; n++ {
		fn := 4*(2*n+1)*(2*n+1) - 12*n + f[n-1]
		f[n] = fn
	}

	fmt.Println(f[500], time.Since(start)) // 669171001, 68.82µs
}
