package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	start := time.Now()
	first := big.NewInt(1)
	second := big.NewInt(1)

	for i := 2; ; i++ {
		if len(second.Text(10)) >= 1000 {
			fmt.Println(i, time.Since(start)) // 4782, 21.89537ms
			break
		}

		first.Add(first, second)
		first, second = second, first
	}
}
