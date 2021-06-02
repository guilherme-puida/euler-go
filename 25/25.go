package main

import (
	"fmt"
	"math/big"
)

func main() {
	first := big.NewInt(1)
	second := big.NewInt(1)

	for i := 2; ; i++ {
		if len(second.Text(10)) >= 1000 {
			fmt.Println(i) // 4782
			break
		}

		first.Add(first, second)
		first, second = second, first
	}
}
