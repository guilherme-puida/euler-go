package main

import (
	"fmt"
	"math/big"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	var n big.Int
	n.MulRange(1, 100)
	sol := 0

	for _, e := range n.Text(10) {
		a, _ := strconv.Atoi(string(e))
		sol += a
	}

	fmt.Println(sol, time.Since(start)) // 648, 18.408µs
}
