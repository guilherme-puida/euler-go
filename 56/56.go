package main

import (
	"fmt"
	"math/big"
	"strconv"
	"time"
)

func SumDigits(a, b int) int {
	var n big.Int
	n.Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), nil)
	sum := 0
	for _, e := range n.Text(10) {
		digit, _ := strconv.Atoi(string(e))
		sum += digit
	}

	return sum
}

func main() {
	start := time.Now()
	maxSum := 0
	for a := 1; a < 100; a++ {
		for b := 1; b < 100; b++ {
			curSum := SumDigits(a, b)
			if curSum > maxSum {
				maxSum = curSum
			}
		}
	}

	fmt.Println(maxSum, time.Since(start)) // 972, 22.996605ms
}
