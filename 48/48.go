package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	start := time.Now()
	var sum big.Int
	for i := 1; i <= 1000; i++ {
		var n big.Int
		i64 := int64(i)
		n.Exp(big.NewInt(i64), big.NewInt(i64), nil)
		sum.Add(&sum, &n)
	}

	s := sum.Text(10)
	fmt.Println(s[len(s)-10:], time.Since(start)) // 9110846700, 2.463817ms
}
