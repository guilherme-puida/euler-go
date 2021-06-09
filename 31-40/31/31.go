package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	coins := [...]int{1, 2, 5, 10, 20, 50, 100, 200}
	var sol [201]int
	sol[0] = 1

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= 200; j++ {
			sol[j] += sol[j-coins[i]]
		}
	}

	fmt.Println(sol[len(sol)-1], time.Since(start)) // 73682, 1.182µs

}
